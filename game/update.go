package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update handles the game logic, like jumping, entity movement, and collision detection
func (g *Game) Update() error {
	inputState := g.input.Update()
	g.character.Input(inputState)

	if g.world.GameOver() {
		if !g.died {
			g.world.Die()
			g.character.Die()
			g.died = true
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if resetErr := g.ResetGame(); resetErr != nil {
				return resetErr
			}
		}
		return nil
	}
	g.world.Update()
	return nil
}
