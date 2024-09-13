package entity

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
}

type BaseEntity struct {
	width       float64
	height      float64
	scaleFactor float64
	frames      []*ebiten.Image
	totalFrames int
	frameIndex  int
	frameCount  int
	frameDelay  int
	kind        Kind
	xPos        float64
	yPos        float64
}

func New(img *ebiten.Image, frameWidth, frameHeight, frameCount, frameDelay, frameRow int, kind Kind, scaleFactor float64) BaseEntity {
	width := float64(frameWidth) * scaleFactor
	height := float64(frameHeight) * scaleFactor

	rowHeight := frameRow * frameHeight
	frames := make([]*ebiten.Image, frameCount)
	for i := 0; i < frameCount; i++ {
		frame, ok := img.SubImage(image.Rect(i*frameWidth, rowHeight, (i+1)*frameWidth, frameHeight+rowHeight)).(*ebiten.Image)
		if !ok {
			panic("failed to create frame for entity")
		}
		frames[i] = frame
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
	screen.DrawImage(e.frames[e.frameIndex], op)

	vector.DrawFilledRect(screen, float32(e.xPos), float32(e.yPos), float32(e.width), float32(e.height), color.RGBA{R: 255}, false)
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
