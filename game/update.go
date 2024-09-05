package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	MinObstacleGap     = 250 // Minimum gap between obstacles
	MaxObstacleGap     = 400 // Maximum gap between obstacles
	LevelJumpThreshold = 50  // Number of successful jumps to advance to the next level

)

// Update handles the game logic, like jumping, obstacle movement, and collision detection
func (g *Game) Update() error {
	// Scroll the background to the left
	g.BackgroundX -= g.BackgroundSpeed

	// Reset the background position for a seamless loop
	if g.BackgroundX <= -ScreenWidth {
		g.BackgroundX = 0
	}

	if g.GameOver {
		// Restart game on space bar press
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.ResetGame()
		}
		return nil
	}

	// If we're in the level greeting phase, manage the countdown
	if g.InLevelGreeting {
		g.HandleCountdown()
		return nil
	}

	// Handle jump mechanics
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !g.IsJumping {
		g.VelocityY = -12
		g.IsJumping = true
	}

	// Apply gravity and update dino's position
	if g.IsJumping {
		g.DinoY += g.VelocityY
		g.VelocityY += g.Gravity

		// Stop falling when reaching the ground
		if g.DinoY >= float64(groundY-DinoHeight) {
			g.DinoY = float64(groundY - DinoHeight)
			g.IsJumping = false
			g.VelocityY = 0
		}
	}

	g.UpdateClouds()

	// Update clouds' positions
	for i := range g.Clouds {
		g.Clouds[i].X -= g.Clouds[i].Speed
		if g.Clouds[i].X < -CloudWidth { // Reset cloud position when it goes off-screen
			g.Clouds[i].X = ScreenWidth
			g.Clouds[i].Y = g.RNG.Float64() * (ScreenHeight / 2) // Reposition cloud at random height
		}
	}

	// Update obstacles' positions
	for i := range g.Obstacles {
		g.Obstacles[i].X -= g.Obstacles[i].Speed
	}

	// Remove obstacles that move off-screen and spawn new ones
	g.Obstacles = g.FilterObstacles()

	// Ensure proper gap between consecutive obstacles
	if len(g.Obstacles) == 0 || (g.Obstacles[len(g.Obstacles)-1].X < ScreenWidth-float64(MinObstacleGap)) {
		g.AddObstacleWithGap()
	}

	// Track jumps and score
	g.TrackJumpsAndScore()

	// Check collisions with each obstacle
	for _, obs := range g.Obstacles {
		if g.CollisionDetected(obs) {
			g.GameOver = true
		}
	}

	return nil
}

// InitializeClouds creates initial clouds at random positions
func (g *Game) InitializeClouds() {
	for i := 0; i < 5; i++ { // Create 5 initial clouds
		g.AddCloud()
	}
}

// AddCloud adds a new cloud at a random height and speed
func (g *Game) AddCloud() {
	cloud := Cloud{
		X:     ScreenWidth,                                                   // Start the cloud on the right side of the screen
		Y:     MinCloudY + g.RNG.Float64()*(MaxCloudY-MinCloudY),             // Restrict cloud Y position to upper part of screen
		Speed: MinCloudSpeed + g.RNG.Float64()*(MaxCloudSpeed-MinCloudSpeed), // Random speed in the range
	}
	g.Clouds = append(g.Clouds, cloud)
}

// UpdateClouds moves the clouds and resets them when they go off-screen
func (g *Game) UpdateClouds() {
	for i := range g.Clouds {
		g.Clouds[i].X -= g.Clouds[i].Speed // Move the cloud to the left
		if g.Clouds[i].X < -CloudWidth {   // If the cloud goes off-screen
			g.Clouds[i].X = ScreenWidth                          // Reposition it to the right
			g.Clouds[i].Y = g.RNG.Float64() * (ScreenHeight / 2) // Random height
		}
	}
}

// HandleCountdown manages the countdown before each level
func (g *Game) HandleCountdown() {
	elapsed := time.Since(g.CountdownStart).Seconds()

	// Handle fade-in/out based on elapsed time
	if elapsed > 1 {
		g.Countdown--
		g.CountdownAlpha = 1.0 // Reset alpha for new countdown number
		g.CountdownStart = time.Now()
	}

	if g.Countdown < 0 {
		// Countdown finished, start the level
		g.InLevelGreeting = false
		g.ObstacleSpeed += 1.0 // Increase difficulty by increasing obstacle speed
	}
}

// TrackJumpsAndScore checks if the dino has successfully jumped over an obstacle
func (g *Game) TrackJumpsAndScore() {
	for i := range g.Obstacles {
		// Check if the obstacle has completely passed the dino (X + width is less than dino position)
		if g.Obstacles[i].X+ObstacleWidth < 40 && !g.Obstacles[i].Passed {
			g.Obstacles[i].Passed = true // Mark the obstacle as passed
			g.Jumps++                    // Increment the jump count
			g.Score += 10                // Increment the score by 10

			// Check if the player needs to level up
			g.HandleLevelProgression()
		}
	}
}

// HandleLevelProgression checks if the dino should level up based on the number of jumps
func (g *Game) HandleLevelProgression() {
	if g.Jumps >= LevelJumpThreshold {
		// Move to the next level
		g.Level++
		g.Jumps = 0 // Reset the jump counter for the next level
		g.InLevelGreeting = true
		g.Countdown = 3 // Start countdown for new level
		g.CountdownAlpha = 1.0
		g.CountdownStart = time.Now()

		// Increase difficulty by increasing the obstacle speed
		g.ObstacleSpeed += 1.0
	}
}

// AddObstacleWithGap adds a new obstacle ensuring there's a proper gap between obstacles
func (g *Game) AddObstacleWithGap() {
	lastObstacle := g.Obstacles[len(g.Obstacles)-1]

	// Create a new obstacle with a random gap between Min and Max values
	gap := g.RNG.Float64()*(MaxObstacleGap-MinObstacleGap) + MinObstacleGap
	newObstacle := Obstacle{
		X:     lastObstacle.X + gap,
		Speed: g.ObstacleSpeed,
	}
	g.Obstacles = append(g.Obstacles, newObstacle)
}

// FilterObstacles removes obstacles that have moved off-screen
func (g *Game) FilterObstacles() []Obstacle {
	var filtered []Obstacle
	for _, obs := range g.Obstacles {
		if obs.X > -ObstacleWidth {
			filtered = append(filtered, obs)
		}
	}
	return filtered
}

// AddObstacle adds a new obstacle to the game
func (g *Game) AddObstacle() {
	newObstacle := Obstacle{
		X:     ScreenWidth + g.RNG.Float64()*200, // Random starting X position
		Speed: g.ObstacleSpeed,
	}
	g.Obstacles = append(g.Obstacles, newObstacle)
}

// ResetGame resets the game state
func (g *Game) ResetGame() {
	g.DinoY = float64(groundY - DinoHeight)
	g.VelocityY = 0
	g.IsJumping = false
	g.Obstacles = []Obstacle{{X: ScreenWidth, Speed: g.ObstacleSpeed}}
	g.Score = 0
	g.GameOver = false
}

// CollisionDetected checks for a collision between the dino and an obstacle
func (g *Game) CollisionDetected(obs Obstacle) bool {
	dinoRight := 40 + DinoWidth
	dinoBottom := g.DinoY + DinoHeight
	obstacleRight := obs.X + ObstacleWidth
	obstacleTop := groundY - ObstacleHeight

	return dinoRight > int(obs.X) && 40 < obstacleRight && int(dinoBottom) > obstacleTop
}
