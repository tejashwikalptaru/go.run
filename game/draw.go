package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen)
	g.Cloud.Draw(screen)
	g.Level.Draw(screen)
	g.Obstacle.Draw(screen)
	g.Player.Draw(screen)

	// If game over, display message
	if g.GameOver {
		msg := fmt.Sprintf("GAME OVER\nScore: %d\nPress Space to Restart", g.Level.Score())
		text.Draw(screen, msg, g.FontFace, ScreenWidth/4, ScreenHeight/2, color.RGBA{R: 255, A: 255})
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.Level.Score()))
	}
}
