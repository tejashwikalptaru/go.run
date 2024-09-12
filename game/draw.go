package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	// If game over, display message
	if g.GameOver {
		msg := fmt.Sprintf("GAME OVER\n\nScore: 100\nPress Space to Restart")
		op := &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				LineSpacing:  50,
				PrimaryAlign: text.AlignCenter,
			},
		}
		op.GeoM.Translate(ScreenWidth/2, ScreenHeight/6)
		op.ColorScale.ScaleWithColor(color.RGBA{R: 255, A: 255})
		text.Draw(screen, msg, &text.GoTextFace{
			Source: g.textFaceSource,
			Size:   48,
		}, op)
		return
	}
	g.world.Draw(screen)
}
