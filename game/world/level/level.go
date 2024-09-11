package level

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/background"
	"github.com/tejashwikalptaru/go.run/game/obstacle"
	"math/rand"
	"time"
)

type Level struct {
	screenWidth, screenHeight      float64
	parallax                       *background.Parallax
	obstacles                      []*obstacle.Obstacle
	obstacleSpeed                  float64
	minObstacleGap, maxObstacleGap float64
	rng                            *rand.Rand
}

func NewLevel(screenWidth, screenHeight float64, parallax *background.Parallax, obstacles []obstacle.Obstacle) *Level {
	level := &Level{
		screenWidth:    screenWidth,
		screenHeight:   screenHeight,
		parallax:       parallax,
		obstacleSpeed:  5,
		minObstacleGap: 250,
		maxObstacleGap: 400,
		rng:            rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	level.distribute(obstacles)
	return level
}

func (l *Level) Update() {
	l.parallax.Update()
	for i := len(l.obstacles) - 1; i >= 0; i-- {
		if l.obstacles[i].XPosition() < -l.obstacles[i].Width() {
			l.obstacles = append(l.obstacles[:i], l.obstacles[i+1:]...)
			continue
		}
		l.obstacles[i].SetXPosition(l.obstacles[i].XPosition() - l.obstacleSpeed)
		l.obstacles[i].Update()
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	l.parallax.Draw(screen)
	for i := range l.obstacles {
		l.obstacles[i].Draw(screen)
	}
}

func (l *Level) distribute(obstacles []obstacle.Obstacle) {
	if obstacles == nil || len(obstacles) == 0 {
		return
	}

	const totalObstacles = 50
	groundY := l.screenHeight - 40
	lastXPos := l.screenWidth + 300

	rand.Shuffle(len(obstacles), func(i, j int) {
		obstacles[i], obstacles[j] = obstacles[j], obstacles[i]
	})

	inAirOffset := groundY - 50
	groundOffset := groundY

	for i := 0; i < totalObstacles; i++ {
		obs := obstacles[l.rng.Intn(len(obstacles))]

		if i != 0 {
			gap := l.rng.Float64()*(l.maxObstacleGap-l.minObstacleGap) + l.minObstacleGap
			lastXPos += gap
		}
		obs.SetXPosition(lastXPos)

		switch obs.Kind() {
		case obstacle.ObstacleKindGround:
			obs.SetYPosition(groundOffset - obs.Height())
		case obstacle.ObstacleKindInAir:
			obs.SetYPosition(inAirOffset - obs.Height())
		case obstacle.ObstacleKindRandom:
			randomOffset := l.rng.Float64()*(groundOffset-inAirOffset) + inAirOffset
			obs.SetYPosition(randomOffset - obs.Height())
		default:
			panic("unknown obstacle")
		}
		l.obstacles = append(l.obstacles, &obs)
	}
}

// Clear returns true if the level is cleared by player
func (l *Level) Clear() bool {
	return len(l.obstacles) == 0 // if no obstacle left on screen
}
