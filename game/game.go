package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/tejashwikalptaru/go.run/resources/fonts"

	"github.com/tejashwikalptaru/go.run/game/background"
	"github.com/tejashwikalptaru/go.run/game/character"
	"github.com/tejashwikalptaru/go.run/game/enemy"
	"github.com/tejashwikalptaru/go.run/game/music"
	"github.com/tejashwikalptaru/go.run/game/stage"
)

const (
	ScreenWidth    = 800
	ScreenHeight   = 400
	LevelThreshold = 5
)

// Game struct holds game state variables
type Game struct {
	RNG            *rand.Rand
	Scene          *background.Scene
	Cloud          *background.Cloud
	Obstacle       *enemy.Obstacle
	Player         *character.Player
	Level          *stage.Level
	MusicManager   *music.Manager
	TextFaceSource *text.GoTextFaceSource
	GameOver       bool
	debug          bool
}

// Layout defines the screen dimensions
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// NewGame initializes a new game instance
func NewGame(debug bool) (*Game, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	textFaceSource, textFaceSourceErr := fonts.LoadFont(fonts.ManaSpace)
	if textFaceSourceErr != nil {
		return nil, textFaceSourceErr
	}

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
	cloud, cloudErr := background.NewCloud(ScreenWidth, ScreenHeight, rng)
	if cloudErr != nil {
		return nil, cloudErr
	}

	// initialise character
	player, playerErr := character.NewPlayer(ScreenWidth, scene.GroundY(), musicManager, debug)
	if playerErr != nil {
		return nil, playerErr
	}

	// initialise obstacle
	obstacle, obstacleErr := enemy.NewObstacle(ScreenWidth, scene.GroundY(), player, rng, LevelThreshold, debug)
	if obstacleErr != nil {
		return nil, obstacleErr
	}

	game := &Game{
		TextFaceSource: textFaceSource,
		RNG:            rng,
		Scene:          scene,
		Cloud:          cloud,
		Obstacle:       obstacle,
		Player:         player,
		GameOver:       false,
		debug:          debug,
		MusicManager:   musicManager,
	}

	// initialise stage
	level := stage.NewLevel(ScreenWidth, ScreenHeight, textFaceSource, LevelThreshold)
	game.Level = level

	// Start playing the background music
	game.MusicManager.PlayBackground()

	return game, nil
}

// ResetGame resets the game state
func (g *Game) ResetGame() error {
	g.Scene.Reset()
	g.Obstacle.Reset()
	g.Player.Reset()
	g.Level = stage.NewLevel(ScreenWidth, ScreenHeight, g.TextFaceSource, LevelThreshold)
	g.GameOver = false
	return nil
}
