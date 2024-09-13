package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/entity"
)

type Player struct {
	entity.BaseEntity
}

func New(screenWidth, groundY float64, img *ebiten.Image, frameWidth, frameHeight, frameCount int) *Player {
	player := entity.New(img, frameWidth, frameHeight, frameCount, 5, "player", 2.0)
	return &Player{player}
}
