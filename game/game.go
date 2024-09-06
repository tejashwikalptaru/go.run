package game

import (
	"math/rand"
	"time"

	"github.com/tejashwikalptaru/go.run/game/background"
	"github.com/tejashwikalptaru/go.run/game/character"
	"github.com/tejashwikalptaru/go.run/game/enemy"
	"github.com/tejashwikalptaru/go.run/game/music"
	"github.com/tejashwikalptaru/go.run/game/stage"
	"golang.org/x/image/font"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 400
)

// Game struct holds game state variables
type Game struct {
	RNG          *rand.Rand
	Scene        *background.Scene
	Cloud        *background.Cloud
	Obstacle     *enemy.Obstacle
	Player       *character.Player
	Level        *stage.Level
	MusicManager *music.Manager
	FontFace     font.Face
	GameOver     bool
}

// Layout defines the screen dimensions
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// NewGame initializes a new game instance
func NewGame(fontFace font.Face) (*Game, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// initialise music manager
	musicManager, musicErr := music.NewMusicManager()
	if musicErr != nil {
		return nil, musicErr
	}

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
	player, playerErr := character.NewPlayer(scene.GroundY(), musicManager)
	if playerErr != nil {
		return nil, playerErr
	}

	// initialise obstacle
	obstacle, obstacleErr := enemy.NewObstacle(ScreenWidth, scene.GroundY(), player, rng)
	if obstacleErr != nil {
		return nil, obstacleErr
	}

	game := &Game{
		FontFace:     fontFace,
		RNG:          rng,
		Scene:        scene,
		Cloud:        cloud,
		Obstacle:     obstacle,
		Player:       player,
		GameOver:     false,
		MusicManager: musicManager,
	}

	// initialise stage
	level := stage.NewLevel(ScreenWidth, ScreenHeight, obstacle.IncreaseSpeed, obstacle.Reset, fontFace)
	game.Level = level

	// Start playing the background music
	game.MusicManager.PlayBackground()

	return game, nil
}

// ResetGame resets the game state
func (g *Game) ResetGame() error {
	g.Obstacle.Reset()
	g.Level = stage.NewLevel(ScreenWidth, ScreenHeight, g.Obstacle.IncreaseSpeed, g.Obstacle.Reset, g.FontFace)
	g.GameOver = false
	return nil
}
