package stage

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tejashwikalptaru/go.run/game/entity"
	"github.com/tejashwikalptaru/go.run/game/music"
	"github.com/tejashwikalptaru/go.run/game/world/level"
	"github.com/tejashwikalptaru/go.run/resource"
)

type imageData struct {
	*ebiten.Image
	scaleX, scaleY float64
	width, height  float64
	alpha          float32
	fadeIn         bool
}

type textData struct {
	textSize          float64
	textAlpha         float32
	textAlphaDelta    float32
	fontFace          *text.GoTextFaceSource
	text              string
	welcomeTextWidth  float64
	welcomeTextHeight float64
}

type Splash struct {
	splashImg *imageData
	engineImg *imageData

	screenWidth, screenHeight float64
	started                   bool

	// music
	music *music.Manager

	// text
	welcomeText     *textData
	developedByText *textData

	done           bool
	startKbdListen bool
}

const fadeSpeed = 0.01

func NewSplash(screenWidth, screenHeight float64) Stage {
	splash := resource.Provider{}.Image("images/splash.jpg")
	splashWidth, splashHeight := float64(splash.Bounds().Dx()), float64(splash.Bounds().Dy())

	engine := resource.Provider{}.Image("images/engine.png")
	engineWidth, engineHeight := float64(engine.Bounds().Dx()), float64(engine.Bounds().Dy())

	fontFace := resource.Provider{}.TextFaceSource("fonts/JungleAdventurer.ttf")
	welcomeText := "Press space to start"
	textSize := 36.0
	welcomeTextWidth, welcomeTextHeight := text.Measure(welcomeText, &text.GoTextFace{
		Source: fontFace,
		Size:   textSize,
	}, 0)
	welcomeTextData := &textData{
		textSize:          textSize,
		textAlpha:         1.0,
		textAlphaDelta:    0.03,
		fontFace:          fontFace,
		text:              welcomeText,
		welcomeTextWidth:  welcomeTextWidth,
		welcomeTextHeight: welcomeTextHeight,
	}

	developedBy := "Tejashwi Kalp Taru\nhttps://tejashwi.io"
	textSize = 10.0
	developedByTextWidth, developedByTextHeight := text.Measure(developedBy, &text.GoTextFace{
		Source: fontFace,
		Size:   textSize,
	}, 10)
	developedByTextData := &textData{
		textSize:          textSize,
		textAlpha:         1.0,
		fontFace:          fontFace,
		text:              developedBy,
		welcomeTextWidth:  developedByTextWidth,
		welcomeTextHeight: developedByTextHeight,
	}

	return &Splash{
		splashImg: &imageData{
			Image:  splash,
			width:  splashWidth,
			height: splashHeight,
			scaleX: screenWidth / splashWidth,
			scaleY: screenHeight / splashHeight,
			alpha:  0.0,
			fadeIn: false,
		},
		engineImg: &imageData{
			Image:  engine,
			width:  engineWidth,
			height: engineHeight,
			scaleX: screenWidth / engineWidth,
			scaleY: screenHeight / engineHeight,
			alpha:  0.0,
			fadeIn: true,
		},
		screenHeight: screenHeight,
		screenWidth:  screenWidth,

		welcomeText:     welcomeTextData,
		developedByText: developedByTextData,
	}
}

func (s *Splash) Update() {
	if !s.started || s.done {
		return
	}

	if !s.done && s.startKbdListen {
		s.done = inpututil.IsKeyJustPressed(ebiten.KeySpace)
		if s.done {
			s.music.FadeStop()
			return
		}
	}

	if s.engineImg.fadeIn {
		s.engineImg.alpha += fadeSpeed
		if s.engineImg.alpha >= 1.0 {
			s.engineImg.alpha = 1.0
			s.engineImg.fadeIn = false
		}
	}

	if !s.engineImg.fadeIn {
		s.engineImg.alpha -= fadeSpeed
		if s.engineImg.alpha <= 0.0 {
			s.engineImg.alpha = 0.0
			s.splashImg.fadeIn = true
		}
	}

	if s.splashImg.fadeIn {
		s.splashImg.alpha += fadeSpeed
		if s.splashImg.alpha >= 1.0 {
			s.splashImg.alpha = 1.0
			if s.music == nil {
				s.music = music.NewLoopMusic(resource.Provider{}.Reader("music/jungle-story-168459-phantasticbeats.mp3"))
				s.music.Play()
				s.startKbdListen = true
			}
		}

		s.welcomeText.textAlpha += s.welcomeText.textAlphaDelta
		if s.welcomeText.textAlpha >= 1.0 {
			s.welcomeText.textAlpha = 1.0
			s.welcomeText.textAlphaDelta *= -1
		} else if s.welcomeText.textAlpha <= 0.0 {
			s.welcomeText.textAlpha = 0.0
			s.welcomeText.textAlphaDelta *= -1
		}
	}
}

func (s *Splash) Draw(screen *ebiten.Image) {
	if !s.started || s.done {
		return
	}

	op := &ebiten.DrawImageOptions{}

	if s.engineImg.alpha > 0 {
		op.GeoM.Scale(s.engineImg.scaleX, s.engineImg.scaleY)
		op.ColorScale.ScaleAlpha(s.engineImg.alpha)
		screen.DrawImage(s.engineImg.Image, op)
	}

	op.GeoM.Reset()

	if s.splashImg.alpha > 0 {
		op.GeoM.Scale(s.splashImg.scaleX, s.splashImg.scaleY)
		op.ColorScale.ScaleAlpha(s.splashImg.alpha)
		screen.DrawImage(s.splashImg.Image, op)
	}

	if s.splashImg.alpha >= 1.0 {
		stripHeight := s.screenHeight / 10
		vector.DrawFilledRect(screen, 0, float32(s.screenHeight-stripHeight), float32(s.screenWidth), float32(stripHeight), color.RGBA{A: 128}, false)

		textOpts := &text.DrawOptions{}
		textOpts.ColorScale.ScaleAlpha(s.welcomeText.textAlpha)
		textOpts.GeoM.Translate(
			(s.screenWidth-s.welcomeText.welcomeTextWidth)/2,
			s.screenHeight-stripHeight+(stripHeight-s.welcomeText.welcomeTextHeight)/2)
		text.Draw(screen, s.welcomeText.text, &text.GoTextFace{
			Source: s.welcomeText.fontFace,
			Size:   s.welcomeText.textSize,
		}, textOpts)

		textOpts = &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				PrimaryAlign: text.AlignStart,
				LineSpacing:  10,
			},
		}
		textOpts.ColorScale.ScaleAlpha(s.developedByText.textAlpha)
		textOpts.GeoM.Translate(s.screenWidth-s.developedByText.welcomeTextWidth-10, s.screenHeight-stripHeight+(stripHeight-s.welcomeText.welcomeTextHeight)+10)
		text.Draw(screen, s.developedByText.text, &text.GoTextFace{
			Source: s.developedByText.fontFace,
			Size:   s.developedByText.textSize,
		}, textOpts)
	}
}

func (s *Splash) NextLevel() bool {
	return false // indicate no level left in this stage
}

func (s *Splash) CurrentLevel() (int, *level.Level) {
	return 0, nil
}

func (s *Splash) Name() string {
	return "Welcome"
}

func (s *Splash) Kind() Kind {
	return KindSplash
}

func (s *Splash) Begin() {
	if s.started {
		return
	}
	s.started = true
}

func (s *Splash) LevelClear() bool {
	return s.done
}

func (s *Splash) CheckCollision(entity *entity.BaseEntity) (bool, entity.Kind) {
	return false, ""
}

func (s *Splash) Die() {

}
