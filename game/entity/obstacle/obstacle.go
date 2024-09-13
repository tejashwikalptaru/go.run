package obstacle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/entity"
)

const (
	KindGround entity.Kind = "ground"
	KindInAir  entity.Kind = "in_air"
	KindRandom entity.Kind = "random"
)

type Obstacle struct {
	entity.BaseEntity
}

func New(img *ebiten.Image, frameWidth, frameHeight, frameCount int, obstacleType entity.Kind) *Obstacle {
	obstacle := entity.New(img, frameWidth, frameHeight, frameCount, 5, 0, obstacleType, 1.5)

	return &Obstacle{obstacle}
}
