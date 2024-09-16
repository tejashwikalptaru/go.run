package stage

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/tejashwikalptaru/go.run/game/entity"
	"github.com/tejashwikalptaru/go.run/game/world/level"
)

type gameStage struct {
	levels                    []*level.Level
	name                      string
	currentLevel              int
	stageClear                bool
	greetingDone              bool
	screenWidth, screenHeight float64
	textFaceSource            *text.GoTextFaceSource
}

func NewStage(name string, screenWidth, screenHeight float64, textFaceSource *text.GoTextFaceSource, levels []*level.Level) Stage {
	return &gameStage{
		levels:         levels,
		name:           name,
		currentLevel:   0,
		screenWidth:    screenWidth,
		screenHeight:   screenHeight,
		textFaceSource: textFaceSource,
	}
}

func (s *gameStage) Update() {
	s.levels[s.currentLevel].Update()
}

func (s *gameStage) Draw(screen *ebiten.Image) {
	s.levels[s.currentLevel].Draw(screen)
}

func (s *gameStage) NextLevel() bool {
	if s.currentLevel < len(s.levels)-1 {
		s.currentLevel++
		return true
	}
	return false // indicate no level left in this stage
}

func (s *gameStage) LevelClear() bool {
	if len(s.levels) == 0 {
		return false
	}
	return s.levels[s.currentLevel].Clear()
}

func (s *gameStage) Name() string {
	return s.name
}

func (s *gameStage) Kind() Kind {
	return KindGame
}

func (s *gameStage) Begin() {
	s.levels[s.currentLevel].Begin()
}

func (s *gameStage) CheckCollision(entity *entity.BaseEntity) bool {
	return s.levels[s.currentLevel].CheckCollision(entity)
}
