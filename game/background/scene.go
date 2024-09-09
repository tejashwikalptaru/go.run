package background

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tejashwikalptaru/go.run/resources"
	"github.com/tejashwikalptaru/go.run/resources/images"
)

type backgroundInfo struct {
	image         *ebiten.Image
	width, height int
}

type Scene struct {
	backGroundImage     *backgroundInfo
	newBackGroundImage  *backgroundInfo
	backgrounds         []backgroundInfo
	screenWidth         float64
	backgroundX         float64
	backgroundSpeed     float64
	groundY             float64
	screenHeight        float64
	backgroundIndex     int
	groundHeight        float64
	fadeAlpha           float32
	fadeSpeed           float32
	transitioning       bool
	transitionCompleted bool
}

func NewScene(screenWidth, screenHeight int) (*Scene, error) {
	// Load all background images and store them in a slice
	bgImages := [][]byte{images.BackgroundOne, images.BackgroundTwo, images.BackgroundThree, images.BackgroundFour}
	backgrounds := make([]backgroundInfo, len(bgImages))

	for i, bg := range bgImages {
		image, imageErr := resources.GetImage(bg)
		if imageErr != nil {
			return nil, imageErr
		}
		bgImg := ebiten.NewImageFromImage(image)
		backgrounds[i] = backgroundInfo{
			image:  bgImg,
			width:  bgImg.Bounds().Dx(),
			height: bgImg.Bounds().Dy(),
		}
	}

	groundHeight := 40.0
	return &Scene{
		groundHeight:    groundHeight,
		groundY:         float64(screenHeight) - groundHeight,
		backgroundX:     0.0,
		backgroundSpeed: 1.0,
		screenWidth:     float64(screenWidth),
		screenHeight:    float64(screenHeight),
		backGroundImage: &backgrounds[0],
		backgrounds:     backgrounds,
		backgroundIndex: 0,
		fadeAlpha:       1.0,
		fadeSpeed:       0.01,
	}, nil
}

func (s *Scene) GroundY() float64 {
	return s.groundY
}

func (s *Scene) Reset() {
	s.backgroundIndex = 0
	s.backGroundImage = &s.backgrounds[s.backgroundIndex]
}

// NextScene transitions to the next background image in the sequence
func (s *Scene) NextScene() {
	// Increment the background index
	s.backgroundIndex = (s.backgroundIndex + 1) % len(s.backgrounds)

	// Get the next background image and prepare the transition
	s.newBackGroundImage = &s.backgrounds[s.backgroundIndex]
	s.transitioning = true
	s.transitionCompleted = false
	s.fadeAlpha = 1.0 // Start fade-out process
}

func (s *Scene) Update() {
	// Handle the fading effect if transitioning to a new background
	if s.transitioning {
		if !s.transitionCompleted {
			// Fade out the current background
			s.fadeAlpha -= s.fadeSpeed
			if s.fadeAlpha <= 0 {
				// Fade out completed, mark as fully faded out
				s.backGroundImage = s.newBackGroundImage
				s.newBackGroundImage = nil
				s.fadeAlpha = 0
				s.transitionCompleted = true // Transition to new background
			}
		} else {
			// Fade in the new background
			s.fadeAlpha += s.fadeSpeed
			if s.fadeAlpha >= 1 {
				// Fade in completed, stop the transition
				s.fadeAlpha = 1
				s.transitioning = false // End transition
			}
		}
		return
	}

	// Scroll the background to the left
	s.backgroundX -= s.backgroundSpeed

	// Reset the background position for a seamless loop
	if s.backgroundX <= -s.screenWidth {
		s.backgroundX = 0
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	// Ensure backGroundImage is not nil
	if s.backGroundImage == nil {
		return // Skip drawing if the background image isn't loaded
	}

	// Calculate scaling factors to resize the image to match the window size
	scaleX := s.screenWidth / float64(s.backGroundImage.width)
	scaleY := s.screenHeight / float64(s.backGroundImage.height)

	// Draw the current background image with fade-out effect
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Scale(scaleX, scaleY)
	op1.GeoM.Translate(s.backgroundX, 0)   // Apply translation for the scrolling effect
	op1.ColorScale.ScaleAlpha(s.fadeAlpha) // Fade-out effect applied to current image
	screen.DrawImage(s.backGroundImage.image, op1)

	// Draw the second background image to create the seamless loop
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(scaleX, scaleY)
	op2.GeoM.Translate(s.backgroundX+s.screenWidth, 0) // Second image positioned right after the first
	op2.ColorScale.ScaleAlpha(s.fadeAlpha)             // Fade-out effect applied to current image
	screen.DrawImage(s.backGroundImage.image, op2)

	// Draw base ground
	vector.DrawFilledRect(screen, 0, float32(s.groundY), float32(s.screenWidth), float32(s.screenHeight), color.RGBA{R: 139, G: 69, B: 19, A: 255}, false)
}
