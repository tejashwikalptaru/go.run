package world

import "github.com/hajimehoshi/ebiten/v2"

type Stage struct {
	levels       []*Level
	name         string
	currentLevel int
	stageClear   bool
}

func NewStage(name string, levels []*Level) *Stage {
	return &Stage{
		levels:       levels,
		name:         name,
		currentLevel: 0,
	}
}

func (s *Stage) Update() {
	s.levels[s.currentLevel].Update()
}

func (s *Stage) Draw(screen *ebiten.Image) {
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
