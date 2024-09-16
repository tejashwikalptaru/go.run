package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type Kind string

type Entity interface {
	Update()
	Draw(screen *ebiten.Image)
	SetXPosition(xPosition float64)
	XPosition() float64
	SetYPosition(yPosition float64)
	YPosition() float64
	Kind() Kind
	Width() float64
	Height() float64
	ScaleFactor() float64
	CollidesWith(other *BaseEntity) bool
	BoundingBox() image.Rectangle
}

type Vec2d struct {
	X float64
	Y float64
}

type Frame struct {
	Image            *ebiten.Image
	TightBoundingBox *image.Rectangle
	Outline          []Vec2d
	DataComputed     bool
}

type BaseEntity struct {
	width       float64
	height      float64
	scaleFactor float64
	frames      []Frame
	totalFrames int
	frameIndex  int
	frameCount  int
	frameDelay  int
	kind        Kind
	xPos        float64
	yPos        float64
}

func New(img *ebiten.Image, frameWidth, frameHeight, frameCount, frameDelay, frameRow int, kind Kind, scaleFactor float64) BaseEntity {
	width := float64(frameWidth)
	height := float64(frameHeight)

	rowHeight := frameRow * frameHeight
	frames := make([]Frame, frameCount)
	for i := 0; i < frameCount; i++ {
		frame, ok := img.SubImage(image.Rect(i*frameWidth, rowHeight, (i+1)*frameWidth, frameHeight+rowHeight)).(*ebiten.Image)
		if !ok {
			panic("failed to create frame for entity")
		}
		frames[i] = Frame{
			Image: frame,
		}
	}
	return BaseEntity{
		width:       width,
		height:      height,
		scaleFactor: scaleFactor,
		frames:      frames,
		totalFrames: len(frames),
		frameDelay:  frameDelay,
		kind:        kind,
	}
}

func (e *BaseEntity) Update() {
	frame := &e.frames[e.frameIndex]
	if !frame.DataComputed {
		rect := computeTightBoundingBox(frame.Image)
		frame.TightBoundingBox = &rect

		outline := generateCollisionOutline(frame.Image)
		simplifiedOutline := simplifyOutline(outline, 10)
		frame.Outline = simplifiedOutline

		frame.DataComputed = true
	}
	e.frameCount++
	if e.frameCount >= e.frameDelay {
		e.frameIndex = (e.frameIndex + 1) % e.totalFrames
		e.frameCount = 0
	}
}

func (e *BaseEntity) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(e.scaleFactor, e.scaleFactor)
	op.GeoM.Translate(e.xPos, e.yPos)
	screen.DrawImage(e.frames[e.frameIndex].Image, op)

	if e.frames[e.frameIndex].DataComputed {
		eTransformedOutline := transformOutline(
			e.frames[e.frameIndex].Outline,
			Vec2d{X: e.XPosition(), Y: e.YPosition()},
			e.scaleFactor,
			0.0, // Replace with e.rotation if applicable
		)
		drawOutline(screen, eTransformedOutline, color.RGBA{G: 255, A: 255})

		rect := e.BoundingBox()
		x := float32(rect.Min.X)
		y := float32(rect.Min.Y)
		width := float32(rect.Max.X - rect.Min.X)
		height := float32(rect.Max.Y - rect.Min.Y)
		vector.DrawFilledRect(screen, x, y, width, height, color.RGBA{R: 255}, false)
	}
}

func (e *BaseEntity) SetXPosition(xPosition float64) {
	e.xPos = xPosition
}

func (e *BaseEntity) XPosition() float64 {
	return e.xPos
}

func (e *BaseEntity) SetYPosition(yPosition float64) {
	e.yPos = yPosition
}

func (e *BaseEntity) YPosition() float64 {
	return e.yPos
}

func (e *BaseEntity) Kind() Kind {
	return e.kind
}

func (e *BaseEntity) Width() float64 {
	return e.width
}

func (e *BaseEntity) Height() float64 {
	return e.height
}

func (e *BaseEntity) ScaleFactor() float64 {
	return e.scaleFactor
}

func (e *BaseEntity) BoundingBox() image.Rectangle {
	frame := &e.frames[e.frameIndex]

	bbox := frame.TightBoundingBox

	// Apply scaling
	scaledWidth := float64(bbox.Dx()) * e.scaleFactor
	scaledHeight := float64(bbox.Dy()) * e.scaleFactor

	// Calculate the top-left corner position, considering any offset from the tight bounding box
	x := e.xPos + float64(bbox.Min.X)*e.scaleFactor
	y := e.yPos + float64(bbox.Min.Y)*e.scaleFactor

	return image.Rect(
		int(x),
		int(y),
		int(x+scaledWidth),
		int(y+scaledHeight),
	)
}

func (e *BaseEntity) CollidesWith(other *BaseEntity) bool {
	if !e.frames[e.frameIndex].DataComputed || !other.frames[other.frameIndex].DataComputed {
		return false
	}

	if !e.BoundingBox().Overlaps(other.BoundingBox()) {
		return false
	}

	// Step 2: Transform Outlines
	eTransformedOutline := transformOutline(
		e.frames[e.frameIndex].Outline,
		Vec2d{X: e.XPosition(), Y: e.YPosition()},
		e.scaleFactor,
		0.0, // Replace with e.rotation if applicable
	)
	otherTransformedOutline := transformOutline(
		other.frames[other.frameIndex].Outline,
		Vec2d{X: other.XPosition(), Y: other.YPosition()},
		other.scaleFactor,
		0.0, // Replace with other.rotation if applicable
	)

	// Step 3: Polygon Collision Detection using SAT
	return polygonsCollide(eTransformedOutline, otherTransformedOutline)
}
