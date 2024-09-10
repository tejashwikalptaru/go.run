package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Obstacle struct {
	frames          []*ebiten.Image
	width           float64
	height          float64
	collisionTop    float64
	collisionLeft   float64
	collisionWidth  float64
	collisionHeight float64
}

func NewObstacle(img ebiten.Image, frameWidth, frameHeight, frameCount int) *Obstacle {
	frames := make([]*ebiten.Image, frameCount)
	for i := 0; i < frameCount; i++ {
		frame, ok := img.SubImage(image.Rect(i*frameWidth, 0, (i+1)*frameWidth, frameHeight)).(*ebiten.Image)
		if !ok {
			panic("failed to create frame for obstacle")
		}
		frames[i] = frame
	}
	return &Obstacle{
		frames:          frames,
		width:           float64(frameWidth),
		height:          float64(frameHeight),
		collisionTop:    0,
		collisionLeft:   0,
		collisionWidth:  float64(frameWidth),
		collisionHeight: float64(frameHeight),
	}
}

func (o *Obstacle) Update() {

}

func (o *Obstacle) Draw(screen *ebiten.Image) {

}
