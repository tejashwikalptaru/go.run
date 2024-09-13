package obstacle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Kind string

const (
	KindGround Kind = "ground"
	KindInAir  Kind = "in_air"
	KindRandom Kind = "random"
)

type Obstacle struct {
	frames      []*ebiten.Image
	totalFrames int
	frameIndex  int
	frameCount  int
	frameDelay  int
	width       float64
	height      float64
	scaleFactor float64
	xPos        float64
	yPos        float64
	kind        Kind
}

func New(img *ebiten.Image, frameWidth, frameHeight, frameCount int, obstacleType Kind) *Obstacle {
	width := float64(frameWidth) * 1.5
	height := float64(frameHeight) * 1.5

	frames := make([]*ebiten.Image, frameCount)
	for i := 0; i < frameCount; i++ {
		frame, ok := img.SubImage(image.Rect(i*frameWidth, 0, (i+1)*frameWidth, frameHeight)).(*ebiten.Image)
		if !ok {
			panic("failed to create frame for obstacle")
		}
		frames[i] = frame
	}

	return &Obstacle{
		frames:      frames,
		totalFrames: len(frames),
		frameCount:  frameCount,
		frameDelay:  5,
		width:       width,
		height:      height,
		scaleFactor: 1.5,
		kind:        obstacleType,
	}
}

func (o *Obstacle) Update() {
	o.frameCount++
	if o.frameCount >= o.frameDelay {
		o.frameIndex = (o.frameIndex + 1) % o.totalFrames
		o.frameCount = 0
	}
}

func (o *Obstacle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(o.scaleFactor, o.scaleFactor)
	op.GeoM.Translate(o.xPos, o.yPos)
	screen.DrawImage(o.frames[o.frameIndex], op)
}

func (o *Obstacle) SetXPosition(xPosition float64) {
	o.xPos = xPosition
}

func (o *Obstacle) XPosition() float64 {
	return o.xPos
}

func (o *Obstacle) SetYPosition(yPosition float64) {
	o.yPos = yPosition
}

func (o *Obstacle) YPosition() float64 {
	return o.yPos
}

func (o *Obstacle) Kind() Kind {
	return o.kind
}

func (o *Obstacle) Width() float64 {
	return o.width
}

func (o *Obstacle) Height() float64 {
	return o.height
}
