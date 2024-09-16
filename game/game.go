package game

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/tejashwikalptaru/go.run/game/entity/player"
	"github.com/tejashwikalptaru/go.run/game/world"
	"github.com/tejashwikalptaru/go.run/resource"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 400
)

// Game struct holds game state variables
type Game struct {
	debug          bool
	died           bool
	textFaceSource *text.GoTextFaceSource
	world          *world.World
	character      *player.Player
}

// Layout defines the screen dimensions
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// NewGame initializes a new game instance
func NewGame(debug bool) (*Game, error) {
	textFaceSource := resource.Provider{}.TextFaceSource("fonts/JungleAdventurer.ttf")
	character := player.New(ScreenWidth, ScreenHeight-40)
	game := &Game{
		debug:          debug,
		textFaceSource: textFaceSource,
		character:      character,
		world:          world.New(ScreenWidth, ScreenHeight, textFaceSource, character),
	}
	return game, nil
}

// ResetGame resets the game state
func (g *Game) ResetGame() error {
	g.character = player.New(ScreenWidth, ScreenHeight-40)
	g.world = world.New(ScreenWidth, ScreenHeight, g.textFaceSource, g.character)
	return nil
}
