package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tejashwikalptaru/go-dino/game/character"
	"github.com/tejashwikalptaru/go-dino/game/stage"
	"image/color"
	"math/rand"
)

type obstacleItem struct {
	X      float64
	Speed  float64
	Passed bool
}

type Obstacle struct {
	obstacleWidth  float64
	obstacleHeight float64
	groundY        float64
	minObstacleGap float64
	maxObstacleGap float64
	screenWidth    float64
	obstacleSpeed  float64
	rng            *rand.Rand
	player         *character.Player
	obstacles      []obstacleItem
}

func NewObstacle(screenWidth, groundY float64, player *character.Player, rng *rand.Rand) *Obstacle {
	return &Obstacle{
		obstacleHeight: 40,
		obstacleWidth:  20,
		minObstacleGap: 250,
		maxObstacleGap: 400,
		screenWidth:    screenWidth,
		rng:            rng,
		obstacleSpeed:  5,
		groundY:        groundY,
		player:         player,
		// Start the first obstacle farther away to avoid instant collision
		obstacles: []obstacleItem{{X: screenWidth + 300, Speed: 5}},
	}
}

// IncreaseSpeed increases the speed of the obstacles as the player progresses to new levels
func (o *Obstacle) IncreaseSpeed() {
	o.obstacleSpeed += 1.0
}

// filterObstacles removes obstacles that have moved off-screen
func (o *Obstacle) filterObstacles() []obstacleItem {
	var filtered []obstacleItem
	for _, obs := range o.obstacles {
		if obs.X > -o.obstacleWidth {
			filtered = append(filtered, obs)
		}
	}
	return filtered
}

// addObstacleWithGap adds a new obstacle ensuring there's a proper gap between obstacles
func (o *Obstacle) addObstacleWithGap() {
	lastObstacle := o.obstacles[len(o.obstacles)-1]

	// Create a new obstacle with a random gap between Min and Max values
	gap := o.minObstacleGap + o.rng.Float64()*(o.maxObstacleGap-o.minObstacleGap)
	newObstacle := obstacleItem{
		X:     lastObstacle.X + gap,
		Speed: o.obstacleSpeed,
	}
	o.obstacles = append(o.obstacles, newObstacle)
}

// collisionDetected checks for a collision between the player and an obstacle
func (o *Obstacle) collisionDetected(obs obstacleItem) bool {
	playerRight := 20 + o.player.Width()
	playerBottom := o.player.Y() + o.player.Height()
	obstacleRight := obs.X + o.obstacleWidth
	obstacleTop := o.groundY - o.obstacleHeight

	// Ensure this logic correctly calculates the collision boundaries
	return playerRight > obs.X && 50 < obstacleRight && playerBottom > obstacleTop
}

// TrackJumpsAndScore checks if the player has successfully jumped over an obstacle
func (o *Obstacle) TrackJumpsAndScore(level *stage.Level) bool {
	for i := range o.obstacles {
		// Check if the obstacle has completely passed the player (X + width is less than player's X)
		if o.obstacles[i].X+o.obstacleWidth < 40 && !o.obstacles[i].Passed {
			o.obstacles[i].Passed = true // Mark the obstacle as passed
			level.IncreaseJump()         // Increment jump count
			level.IncreaseScore()        // Increment the score by 10

			// Check for level progression
			return true
		}
	}
	return false
}

// Reset clears the current obstacles and resets the obstacle list
func (o *Obstacle) Reset() {
	o.obstacles = []obstacleItem{{X: o.screenWidth + 300, Speed: o.obstacleSpeed}} // Start with one obstacle far away
}

// Update handles the movement of obstacles and checks for collisions
func (o *Obstacle) Update() bool {
	for i := range o.obstacles {
		o.obstacles[i].X -= o.obstacles[i].Speed // Move the obstacle to the left
	}
	o.obstacles = o.filterObstacles()

	// Ensure proper gap between consecutive obstacles
	if len(o.obstacles) == 0 || (o.obstacles[len(o.obstacles)-1].X < o.screenWidth-o.minObstacleGap) {
		o.addObstacleWithGap()
	}

	// Check collisions with each obstacle
	for _, obs := range o.obstacles {
		if o.collisionDetected(obs) {
			return true // Trigger game over if a collision is detected
		}
	}
	return false
}

// Draw renders the obstacles on the screen
func (o *Obstacle) Draw(screen *ebiten.Image) {
	for _, obs := range o.obstacles {
		vector.DrawFilledRect(screen, float32(obs.X), float32(o.groundY-o.obstacleHeight), float32(o.obstacleWidth), float32(o.obstacleHeight), color.RGBA{R: 255, A: 255}, false)
	}
}
