package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/background"
	"github.com/tejashwikalptaru/go.run/game/enemy"
	"time"
)

type Level struct {
	parallax                    *background.Parallax
	obstacles                   []*enemy.Obstacle
	currentScore, requiredScore int
	ticker                      *time.Ticker
}

func NewLevel(parallax *background.Parallax, obstacles []*enemy.Obstacle) *Level {
	level := &Level{
		parallax:      parallax,
		obstacles:     obstacles,
		requiredScore: 50,
	}
	return level
}

func (l *Level) Update() {
	if l.ticker == nil {
		l.ticker = time.NewTicker(time.Second * 20)
		go func(l *Level) {
			for range l.ticker.C {
				l.currentScore = l.requiredScore
			}
		}(l)
	}
	l.parallax.Update()
	for _, obstacle := range l.obstacles {
		obstacle.Update()
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	l.parallax.Draw(screen)
	for _, obstacle := range l.obstacles {
		obstacle.Draw(screen)
	}
}

func (l *Level) CurrentScore() int {
	return l.currentScore
}

// Clear returns true if the level is cleared by player
func (l *Level) Clear() bool {
	return l.currentScore >= l.requiredScore
}

func (l *Level) Kill() {
	l.ticker.Stop()
}
