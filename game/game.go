package game

import (
	"github.com/tejashwikalptaru/go-dino/game/background"
	"github.com/tejashwikalptaru/go-dino/game/character"
	"github.com/tejashwikalptaru/go-dino/game/enemy"
	"github.com/tejashwikalptaru/go-dino/game/stage"
	"golang.org/x/image/font"
	"math/rand"
	"time"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 400
)

// Game struct holds game state variables
type Game struct {
	FontFace font.Face
	RNG      *rand.Rand
	Scene    *background.Scene
	Cloud    *background.Cloud
	Obstacle *enemy.Obstacle
	Player   *character.Player
	Level    *stage.Level
	GameOver bool
}

// Layout defines the screen dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// NewGame initializes a new game instance
func NewGame(fontFace font.Face) (*Game, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// initialise background scene
	scene, sceneErr := background.NewScene(ScreenWidth, ScreenHeight)
	if sceneErr != nil {
		return nil, sceneErr
	}

	// initialise cloud
	cloud, cloudErr := background.NewCloud(ScreenWidth, ScreenHeight, 5, rng)
	if cloudErr != nil {
		return nil, cloudErr
	}

	// initialise character
	player := character.NewPlayer(scene.GroundY())

	// initialise obstacle
	obstacle := enemy.NewObstacle(ScreenWidth, scene.GroundY(), player, rng)

	game := &Game{
		FontFace: fontFace,
		RNG:      rng,
		Scene:    scene,
		Cloud:    cloud,
		Obstacle: obstacle,
		Player:   player,
		GameOver: false,
	}

	// initialise stage
	level := stage.NewLevel(ScreenWidth, ScreenHeight, game.ResetGame, obstacle.IncreaseSpeed, obstacle.Reset, fontFace)
	game.Level = level
	return game, nil
}

// ResetGame resets the game state
func (g *Game) ResetGame() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	g.Player = character.NewPlayer(g.Scene.GroundY())
	g.Obstacle = enemy.NewObstacle(ScreenWidth, g.Scene.GroundY(), g.Player, rng)
	g.Level = stage.NewLevel(ScreenWidth, ScreenHeight, g.ResetGame, g.Obstacle.IncreaseSpeed, g.Obstacle.Reset, g.FontFace)
	g.GameOver = false // Reset the game over state
}
