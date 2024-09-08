package game

import "github.com/hajimehoshi/ebiten/v2"

// Update handles the game logic, like jumping, obstacle movement, and collision detection
func (g *Game) Update() error {
	// If the game is over, wait for the player to press space to restart
	if g.GameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			// ResetToFirst the game state when space is pressed
			if resetErr := g.ResetGame(); resetErr != nil {
				return resetErr
			}
		}
		// Stop further updates until the game is reset
		return nil
	}

	// Pause game elements during countdown
	if g.Level.IsGreeting() {
		g.Scene.Update() // Update background scene for visual consistency
		g.Cloud.Update() // Continue cloud movement even during countdown
		g.Level.Update() // Handle the countdown
		return nil       // Do not update the player or obstacles during countdown
	}

	// Normal game updates
	g.Scene.Update()
	g.Cloud.Update()
	g.Level.Update()
	g.Player.Update()

	// Check if there is a collision or the obstacle is cleared
	collision, isPowerUpCollision, obstacleCleared := g.Obstacle.Update()
	if collision {
		g.GameOver = !g.Player.IsImmune()
		if !isPowerUpCollision {
			g.MusicManager.PlayCollisionSound()
		}
	}
	// Track jumps and score, and check for level progression
	if obstacleCleared {
		// increase score and jumps count
		g.Level.IncreaseScore()
	}

	if g.Level.Clear() {
		if !g.Player.WalkingToLevelExit() {
			// if walk to level exit is done, transition to next level
			g.Scene.NextScene()
			g.Level.Next()
			g.Obstacle.Prepare() // reset obstacles for next level
			g.Player.Reset()
			g.Obstacle.IncreaseSpeed() // increase obstacle speed
		}
	}

	return nil
}
