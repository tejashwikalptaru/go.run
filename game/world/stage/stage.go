package stage

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/world/level"
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
	CurrentLevel() (int, *level.Level)
	Name() string
	Kind() Kind
	Begin()
}
