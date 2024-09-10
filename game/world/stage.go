package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Stage struct {
	levels                    []*Level
	name                      string
	currentLevel              int
	stageClear                bool
	greetingDone              bool
	screenWidth, screenHeight float64
	textFaceSource            *text.GoTextFaceSource
}

func NewStage(name string, screenWidth, screenHeight float64, textFaceSource *text.GoTextFaceSource, levels []*Level) *Stage {
	return &Stage{
		levels:         levels,
		name:           name,
		currentLevel:   1,
		screenWidth:    screenWidth,
		screenHeight:   screenHeight,
		textFaceSource: textFaceSource,
	}
}

func (s *Stage) Update() {
	s.levels[s.currentLevel].Update()
}

func (s *Stage) Draw(screen *ebiten.Image) {
	//if !s.greetingDone {
	//	msg := fmt.Sprintf("Stage: %s\n", s.name)
	//	op := &text.DrawOptions{}
	//	op.GeoM.Translate(s.screenWidth/3, s.screenHeight/6)
	//	op.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 0, B: 0})
	//	text.Draw(screen, msg, &text.GoTextFace{
	//		Source: s.textFaceSource,
	//		Size:   48,
	//	}, op)
	//	return
	//}
	s.levels[s.currentLevel].Draw(screen)
}

func (s *Stage) NextLevel() {
	if s.currentLevel < len(s.levels)-1 {
		s.levels[s.currentLevel].Kill()
		s.currentLevel++
	}
}

func (s *Stage) CurrentLevel() (int, *Level) {
	return s.currentLevel, s.levels[s.currentLevel]
}

func (s *Stage) Name() string {
	return s.name
}
