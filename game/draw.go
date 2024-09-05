package game

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tejashwikalptaru/go-dino/resources/images"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	cloudImage      *ebiten.Image
	backGroundImage *ebiten.Image
)

func init() {
	cloudImage = ebiten.NewImageFromImage(GetImage(images.Cloud))
	backGroundImage = ebiten.NewImageFromImage(GetImage(images.BackgroundOne))
}

func GetImage(data []byte) image.Image {
	img, _, decodeErr := image.Decode(bytes.NewReader(data))
	if decodeErr != nil {
		log.Fatalf("failed to decode level one_bg image: %v", decodeErr)
	}
	return img
}

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Get the original dimensions of the background image
	backGroundWidth, backGroundHeight := backGroundImage.Size()

	// Calculate scaling factors to resize the image to match the window size
	scaleX := float64(ScreenWidth) / float64(backGroundWidth)
	scaleY := float64(ScreenHeight) / float64(backGroundHeight)

	// Create image drawing options objects for two images (for seamless scrolling)
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Scale(scaleX, scaleY)
	op1.GeoM.Translate(g.BackgroundX, 0) // Apply translation for the scrolling effect
	screen.DrawImage(backGroundImage, op1)

	// Draw the second background image to create the seamless loop
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(scaleX, scaleY)
	op2.GeoM.Translate(g.BackgroundX+float64(ScreenWidth), 0) // Second image positioned right after the first
	screen.DrawImage(backGroundImage, op2)

	for _, cloud := range g.Clouds {
		cloudOp := &ebiten.DrawImageOptions{}
		cloudOp.GeoM.Translate(cloud.X, cloud.Y) // Translate the cloud to its position
		screen.DrawImage(cloudImage, cloudOp)    // Draw the cloud
	}

	// If we're in the level greeting phase, show the greeting and countdown
	if g.InLevelGreeting {
		msg := fmt.Sprintf("Welcome to Level %d", g.Level)
		text.Draw(screen, msg, g.FontFace, ScreenWidth/4, ScreenHeight/3, color.RGBA{R: 255, G: 255, B: 255, A: 255})

		// Countdown logic: Fade-in/out based on alpha value
		countdownColor := color.RGBA{R: 255, G: 0, B: 0, A: uint8(g.CountdownAlpha * 255)}
		countdownText := fmt.Sprintf("%d", g.Countdown)
		text.Draw(screen, countdownText, g.FontFace, ScreenWidth/2, ScreenHeight/2, countdownColor)

		// Update alpha value for smooth fade in/out
		g.CountdownAlpha -= 0.05
		if g.CountdownAlpha < 0 {
			g.CountdownAlpha = 0
		}

		return
	}

	// Draw the ground
	vector.DrawFilledRect(screen, 0, float32(groundY), ScreenWidth, GroundHeight, color.RGBA{R: 139, G: 69, B: 19, A: 255}, false)

	// Draw the dino
	vector.DrawFilledRect(screen, 40, float32(g.DinoY), DinoWidth, DinoHeight, color.RGBA{G: 128, A: 255}, false)

	// Draw each obstacle
	for _, obs := range g.Obstacles {
		vector.DrawFilledRect(screen, float32(obs.X), float32(groundY-ObstacleHeight), float32(ObstacleWidth), float32(ObstacleHeight), color.RGBA{R: 255, A: 255}, false)
	}

	// Draw score or game over message
	if g.GameOver {
		msg := fmt.Sprintf("GAME OVER\nScore: %d\nPress Space to Restart", g.Score)
		text.Draw(screen, msg, g.FontFace, ScreenWidth/4, ScreenHeight/2, color.RGBA{R: 255, A: 255})
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.Score))
	}
}
