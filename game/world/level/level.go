package level

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/tejashwikalptaru/go.run/game/background"
	"github.com/tejashwikalptaru/go.run/game/entity"
	"github.com/tejashwikalptaru/go.run/game/entity/obstacle"
	"github.com/tejashwikalptaru/go.run/game/music"
	"github.com/tejashwikalptaru/go.run/resource"
)

type Level struct {
	screenWidth, screenHeight      float64
	parallax                       *background.Parallax
	obstacles                      []*obstacle.Obstacle
	obstacleSpeed                  float64
	minObstacleGap, maxObstacleGap float64
	rng                            *rand.Rand
	musicManager                   *music.Manager
	levelCompletedMusic            *music.Manager
	started                        bool
	space                          *resolv.Space
	generatedObstacles             int

	providedObstacles []obstacle.Obstacle
}

func NewLevel(screenWidth, screenHeight float64, parallax *background.Parallax, obstacles []obstacle.Obstacle, musicManager *music.Manager) *Level {
	level := &Level{
		screenWidth:         screenWidth,
		screenHeight:        screenHeight,
		parallax:            parallax,
		obstacleSpeed:       5,
		minObstacleGap:      250,
		maxObstacleGap:      400,
		rng:                 rand.New(rand.NewSource(time.Now().UnixNano())),
		musicManager:        musicManager,
		levelCompletedMusic: music.NewMusic(resource.Provider{}.Reader("music/game-level-complete-143022-universfield.mp3")),
		space:               resolv.NewSpace(int(screenWidth), int(screenHeight), 8, 8),
		providedObstacles:   obstacles,
	}
	level.distributeObstacle()
	return level
}

func (l *Level) Update() {
	l.parallax.Update()
	for i := len(l.obstacles) - 1; i >= 0; i-- {
		if l.obstacles[i].XPosition() < -l.obstacles[i].Width() {
			l.obstacles = append(l.obstacles[:i], l.obstacles[i+1:]...)
			if len(l.obstacles) == 0 {
				// all obstacles cleared
				l.levelCompletedMusic.Play()
			}
			continue
		}
		l.obstacles[i].SetXPosition(l.obstacles[i].XPosition() - l.obstacleSpeed)
		l.obstacles[i].Update()
	}
	if len(l.obstacles) > 0 && l.obstacles[len(l.obstacles)-1].XPosition() < l.screenWidth*0.5 {
		l.distributeObstacle()
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	l.parallax.Draw(screen)
	for i := range l.obstacles {
		l.obstacles[i].Draw(screen)
	}
}

func (l *Level) distributeObstacle() {
	if l.providedObstacles == nil || len(l.providedObstacles) == 0 {
		return
	}

	const totalObstacles = 50
	groundY := l.screenHeight - 60

	if l.generatedObstacles >= totalObstacles {
		return
	}

	rand.Shuffle(len(l.providedObstacles), func(i, j int) {
		l.providedObstacles[i], l.providedObstacles[j] = l.providedObstacles[j], l.providedObstacles[i]
	})

	inAirOffset := groundY - 50
	groundOffset := groundY

	var lastXPos float64
	if len(l.obstacles) > 0 {
		lastXPos = l.obstacles[len(l.obstacles)-1].XPosition()
	} else {
		lastXPos = l.screenWidth + 300
	}

	batchSize := min(5, totalObstacles-l.generatedObstacles)

	for i := 0; i < batchSize; i++ {
		obs := l.providedObstacles[l.rng.Intn(len(l.providedObstacles))]

		gap := l.rng.Float64()*(l.maxObstacleGap-l.minObstacleGap) + l.minObstacleGap
		lastXPos += gap
		obs.SetXPosition(lastXPos)

		switch obs.Kind() {
		case obstacle.KindGround:
			obs.SetYPosition(groundOffset - obs.Height())
		case obstacle.KindInAir:
			obs.SetYPosition(inAirOffset - obs.Height())
		case obstacle.KindRandom:
			randomOffset := l.rng.Float64()*(groundOffset-inAirOffset) + inAirOffset
			obs.SetYPosition(randomOffset - obs.Height())
		default:
			panic("unknown entity")
		}

		l.obstacles = append(l.obstacles, &obs)
		l.generatedObstacles++
	}
}

func (l *Level) Clear() bool {
	done := len(l.obstacles) == 0 // if no entity left on screen
	if done {
		l.musicManager.FadeStop()
		return done
	}
	return false // level running
}

func (l *Level) Begin() {
	if l.started {
		return
	}
	l.started = true
	l.musicManager.Play()
}

func (l *Level) CheckCollision(player *entity.BaseEntity) bool {
	for i := range l.obstacles {
		if l.obstacles[i].CollidesWith(player) {
			return true
		}
	}
	return false
}
