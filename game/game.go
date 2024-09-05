package game

import (
	"golang.org/x/image/font"
	"math/rand"
	"time"
)

const (
	ScreenWidth    = 800
	ScreenHeight   = 400
	GroundHeight   = 40
	DinoHeight     = 60
	DinoWidth      = 40
	ObstacleWidth  = 20
	ObstacleHeight = 40
	CloudWidth     = 100
	MinCloudY      = 5                // Minimum height for the clouds (from the top)
	MaxCloudY      = ScreenHeight / 4 // Maximum height for the clouds (upper third of the screen)
	MinCloudSpeed  = 0.5
	MaxCloudSpeed  = 1.0
)

var (
	groundY = ScreenHeight - GroundHeight
)

// Cloud struct represents each moving cloud in the background
type Cloud struct {
	X     float64
	Y     float64
	Speed float64
}

// Obstacle struct represents each obstacle in the game
type Obstacle struct {
	X      float64
	Speed  float64
	Passed bool
}

// Game struct holds game state variables
type Game struct {
	DinoY           float64
	VelocityY       float64
	IsJumping       bool
	Gravity         float64
	Obstacles       []Obstacle // A list of obstacles
	Clouds          []Cloud    // A list of moving clouds
	ObstacleSpeed   float64
	Score           int
	Jumps           int // Count successful jumps over obstacles
	Level           int // Current level
	GameOver        bool
	InLevelGreeting bool      // Whether we are in the level greeting stage
	Countdown       int       // Countdown before the start of each level
	CountdownAlpha  float64   // Alpha value for countdown animation (fade in/out)
	CountdownStart  time.Time // When countdown started
	FontFace        font.Face
	RNG             *rand.Rand
	BackgroundX     float64
	BackgroundSpeed float64
}

// Layout defines the screen dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// NewGame initializes a new game instance
func NewGame(fontFace font.Face) *Game {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	initialObstacles := []Obstacle{{X: ScreenWidth, Speed: 5}}

	game := &Game{
		DinoY:           float64(groundY - DinoHeight),
		VelocityY:       0,
		IsJumping:       false,
		Gravity:         0.6,
		Obstacles:       initialObstacles,
		ObstacleSpeed:   5,
		Score:           0,
		Jumps:           0,
		Level:           1,
		GameOver:        false,
		InLevelGreeting: true, // Start with level greeting
		Countdown:       3,    // Start countdown at 3
		CountdownAlpha:  1.0,  // Start with fully visible countdown
		CountdownStart:  time.Now(),
		FontFace:        fontFace,
		RNG:             rng,
		BackgroundX:     0.0,
		BackgroundSpeed: 1.0,
	}
	game.InitializeClouds() // Initialize clouds at the start
	return game
}
