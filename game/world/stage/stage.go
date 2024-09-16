package stage

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/entity/player"
)

type Kind string

const (
	KindGame   Kind = "Game"
	KindSplash Kind = "Splash"
)

type Stage interface {
	Update()
	Draw(screen *ebiten.Image)
	NextLevel() bool
	LevelClear() bool
	Name() string
	Kind() Kind
	Begin()
	CheckCollision(player *player.Player) bool
}
