package game

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/tejashwikalptaru/go.run/game/character"
	"github.com/tejashwikalptaru/go.run/game/world"
	"github.com/tejashwikalptaru/go.run/resource"
	"math/rand"
	"time"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 400
)

// Game struct holds game state variables
type Game struct {
	RNG            *rand.Rand
	GameOver       bool
	debug          bool
	textFaceSource *text.GoTextFaceSource
	world          *world.World
}

// Layout defines the screen dimensions
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// NewGame initializes a new game instance
func NewGame(debug bool) (*Game, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	game := &Game{
		RNG:            rng,
		GameOver:       false,
		debug:          debug,
		textFaceSource: resource.Provider{}.TextFaceSource("fonts/manaspace/manaspc.ttf"),
		world:          world.New(ScreenWidth, ScreenHeight, character.NewPlayer(ScreenWidth, ScreenHeight-40, nil, false)),
	}
	return game, nil
}

// ResetGame resets the game state
func (g *Game) ResetGame() error {
	g.GameOver = false
	return nil
}
