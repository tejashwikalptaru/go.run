package obstacle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/entity"
)

const (
	KindGround entity.Kind = "obstacle_ground"
	KindInAir  entity.Kind = "obstacle_in_air"
	KindRandom entity.Kind = "obstacle_random"
)

type Obstacle struct {
	entity.BaseEntity
}

func Is(kind entity.Kind) bool {
	return kind == KindGround || kind == KindInAir || kind == KindRandom
}

func New(img *ebiten.Image, frameWidth, frameHeight, frameCount int, obstacleType entity.Kind, scaleFactor float64) *Obstacle {
	obstacle := entity.New(img, frameWidth, frameHeight, frameCount, 5, 0, obstacleType, scaleFactor)

	return &Obstacle{obstacle}
}

func (o *Obstacle) HFlip() *Obstacle {
	o.BaseEntity.HFlip()
	return o
}
