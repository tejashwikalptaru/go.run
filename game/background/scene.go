package background

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tejashwikalptaru/go.run/resources"
	"github.com/tejashwikalptaru/go.run/resources/images"
)

type Scene struct {
	backGroundImage *ebiten.Image
	groundHeight    float64
	groundY         float64
	backgroundX     float64
	backgroundSpeed float64
	screenWidth     float64
	screenHeight    float64
}

func NewScene(screenWidth, screenHeight int) (*Scene, error) {
	image, imageErr := resources.GetImage(images.BackgroundFour)
	if imageErr != nil {
		return nil, imageErr
	}
	groundHeight := 40.0
	return &Scene{
		groundHeight:    groundHeight,
		groundY:         float64(screenHeight) - groundHeight,
		backgroundX:     0.0,
		backgroundSpeed: 1.0,

		screenWidth:     float64(screenWidth),
		screenHeight:    float64(screenHeight),
		backGroundImage: ebiten.NewImageFromImage(image),
	}, nil
}

func (s *Scene) GroundY() float64 {
	return s.groundY
}

func (s *Scene) Update() {
	// Scroll the background to the left
	s.backgroundX -= s.backgroundSpeed

	// ResetToFirst the background position for a seamless loop
	if s.backgroundX <= -s.screenWidth {
		s.backgroundX = 0
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	// Get the original dimensions of the background image
	backGroundWidth, backGroundHeight := s.backGroundImage.Bounds().Dx(), s.backGroundImage.Bounds().Dy()

	// Calculate scaling factors to resize the image to match the window size
	scaleX := s.screenWidth / float64(backGroundWidth)
	scaleY := s.screenHeight / float64(backGroundHeight)

	// Create image drawing options objects for two images (for seamless scrolling)
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Scale(scaleX, scaleY)
	op1.GeoM.Translate(s.backgroundX, 0) // Apply translation for the scrolling effect
	screen.DrawImage(s.backGroundImage, op1)

	// Draw the second background image to create the seamless loop
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(scaleX, scaleY)
	op2.GeoM.Translate(s.backgroundX+s.screenWidth, 0) // Second image positioned right after the first
	screen.DrawImage(s.backGroundImage, op2)

	// Draw base ground
	vector.DrawFilledRect(screen, 0, float32(s.groundY), float32(s.screenWidth), float32(s.screenHeight), color.RGBA{R: 139, G: 69, B: 19, A: 255}, false)
}
