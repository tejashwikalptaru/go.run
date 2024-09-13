package game

import "github.com/hajimehoshi/ebiten/v2"

// Update handles the game logic, like jumping, entity movement, and collision detection
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
	g.world.Update()
	return nil
}
