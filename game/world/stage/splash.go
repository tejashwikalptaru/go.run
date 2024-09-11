package stage

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tejashwikalptaru/go.run/game/music"
	"github.com/tejashwikalptaru/go.run/game/world/level"
	"github.com/tejashwikalptaru/go.run/resource"
	"image/color"
)

type Splash struct {
	img                       *ebiten.Image
	scaleX, scaleY            float64 // Scaling factor
	width, height             float64 // Image height and width
	screenWidth, screenHeight float64
	started                   bool

	// music
	music    *music.Manager
	fontFace *text.GoTextFaceSource
}

func NewSplash(screenWidth, screenHeight float64) Stage {
	img := resource.Provider{}.Image("images/splash.jpg")
	width, height := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())
	return &Splash{
		img:          img,
		width:        width,
		height:       height,
		screenHeight: screenHeight,
		screenWidth:  screenWidth,
		scaleX:       screenWidth / width,
		scaleY:       screenHeight / height,
		fontFace:     resource.Provider{}.TextFaceSource("fonts/JungleAdventurer.ttf"),
	}
}

func (s *Splash) Update() {
}

func (s *Splash) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.scaleX, s.scaleY)
	op.GeoM.Translate(0, 0)
	screen.DrawImage(s.img, op)

	vector.DrawFilledRect(screen, 300, float32(s.screenHeight-10), float32(s.screenWidth), 100, color.RGBA{0, 0, 0, 1}, false)
	top := &text.DrawOptions{}
	top.GeoM.Translate(312, s.screenHeight-50)
	top.ColorScale.ScaleWithColor(color.RGBA{R: 255})
	text.Draw(screen, "Welcome", &text.GoTextFace{
		Source: s.fontFace,
		Size:   48,
	}, top)
}

func (s *Splash) NextLevel() bool {
	panic("implement me")
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
	s.music = music.NewLoopMusic(resource.Provider{}.Reader("music/jungle-story-168459-phantasticbeats.mp3"))
	s.music.Play()
}
