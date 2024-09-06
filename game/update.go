package game

import "github.com/hajimehoshi/ebiten/v2"

// Update handles the game logic, like jumping, obstacle movement, and collision detection
func (g *Game) Update() error {
	// If the game is over, wait for the player to press space to restart
	if g.GameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			// Reset the game state when space is pressed
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

	// Check if there is a collision, triggering game over
	g.GameOver = g.Obstacle.Update()
	if g.GameOver {
		g.MusicManager.PlayCollisionSound()
	}

	// Track jumps and score, and check for level progression
	if g.Obstacle.TrackJumpsAndScore(g.Level) {
		g.Level.HandleLevelProgression()
	}

	return nil
}
