package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/tejashwikalptaru/go.run/resources/fonts"
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
		msg := fmt.Sprintf("GAME OVER\n\nScore: %d\nPress Space to Restart", g.Level.Score())
		op := &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				LineSpacing:  50,
				PrimaryAlign: text.AlignCenter,
			},
		}
		op.GeoM.Translate(ScreenWidth/2, ScreenHeight/6)
		op.ColorScale.ScaleWithColor(color.RGBA{R: 255, A: 255})
		text.Draw(screen, msg, &text.GoTextFace{
			Source: g.TextFaceSource,
			Size:   fonts.DefaultTextSize,
		}, op)
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.Level.Score()))
	}
}
