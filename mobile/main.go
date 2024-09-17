package mobile

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/tejashwikalptaru/go.run/game"
)

func init() {
	// Create game instance
	g, err := game.NewGame(false)
	if err != nil {
		log.Fatal(err)
	}

	mobile.SetGame(g)
}

func Dummy() {}
