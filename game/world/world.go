package world

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tejashwikalptaru/go.run/game/background"
	"github.com/tejashwikalptaru/go.run/game/character"
	"github.com/tejashwikalptaru/go.run/game/obstacle"
	"github.com/tejashwikalptaru/go.run/game/world/level"
	"github.com/tejashwikalptaru/go.run/game/world/stage"
	"github.com/tejashwikalptaru/go.run/resource"
	"image/color"
)

type World struct {
	player       *character.Player
	stages       []stage.Stage
	currentStage int

	fading             bool
	fadeAlpha          float64
	fadeSpeed          float64
	fadeIn             bool
	transitionComplete bool
}

func New(screenWidth, screenHeight float64, textFaceSource *text.GoTextFaceSource, player *character.Player) *World {
	world := World{
		player:       player,
		currentStage: 0,
		fadeAlpha:    0.0,
		fadeSpeed:    0.01,
	}

	splash := stage.NewSplash(screenWidth, screenHeight)
	world.stages = append(world.stages, splash)

	// jungle stage
	jungleStage := stage.NewStage("Jungle", screenWidth, screenHeight, textFaceSource, []*level.Level{
		level.NewLevel(screenWidth, screenHeight, background.NewParallax([]*background.Layer{
			background.NewLayer(screenWidth, screenHeight, 0.1, resource.Provider{}.Image("images/jungle/1/1.png")),
			background.NewLayer(screenWidth, screenHeight, 0.3, resource.Provider{}.Image("images/jungle/1/2.png")),
			background.NewLayer(screenWidth, screenHeight, 0.5, resource.Provider{}.Image("images/jungle/1/3.png")),
			background.NewLayer(screenWidth, screenHeight, 1.0, resource.Provider{}.Image("images/jungle/1/4.png")),
		}), []obstacle.Obstacle{
			*obstacle.New(resource.Provider{}.Image("sprites/enemy/Deceased_walk.png"), 48, 48, 6, obstacle.ObstacleKindGround),
			*obstacle.New(resource.Provider{}.Image("sprites/enemy/Hyena_walk.png"), 48, 48, 6, obstacle.ObstacleKindGround),
			*obstacle.New(resource.Provider{}.Image("sprites/enemy/Mummy_walk.png"), 48, 48, 6, obstacle.ObstacleKindGround),
			*obstacle.New(resource.Provider{}.Image("sprites/enemy/Scorpio_walk.png"), 48, 48, 4, obstacle.ObstacleKindGround),
			*obstacle.New(resource.Provider{}.Image("sprites/enemy/Snake_walk.png"), 48, 48, 4, obstacle.ObstacleKindGround),
			*obstacle.New(resource.Provider{}.Image("sprites/enemy/Vulture_walk.png"), 48, 48, 4, obstacle.ObstacleKindRandom),
		}),
		level.NewLevel(screenWidth, screenHeight, background.NewParallax([]*background.Layer{
			background.NewLayer(screenWidth, screenHeight, 0.1, resource.Provider{}.Image("images/jungle/2/1.png")),
			background.NewLayer(screenWidth, screenHeight, 0.3, resource.Provider{}.Image("images/jungle/2/2.png")),
			background.NewLayer(screenWidth, screenHeight, 0.5, resource.Provider{}.Image("images/jungle/2/3.png")),
			background.NewLayer(screenWidth, screenHeight, 0.7, resource.Provider{}.Image("images/jungle/2/4.png")),
			background.NewLayer(screenWidth, screenHeight, 0.9, resource.Provider{}.Image("images/jungle/2/4.png")),
			background.NewLayer(screenWidth, screenHeight, 1.0, resource.Provider{}.Image("images/jungle/2/6.png")),
		}), nil),
		level.NewLevel(screenWidth, screenHeight, background.NewParallax([]*background.Layer{
			background.NewLayer(screenWidth, screenHeight, 1.0, resource.Provider{}.Image("images/jungle/3/1.png")),
			background.NewLayer(screenWidth, screenHeight, 0.3, resource.Provider{}.Image("images/jungle/3/2.png")),
			background.NewLayer(screenWidth, screenHeight, 0.5, resource.Provider{}.Image("images/jungle/3/3.png")),
			background.NewLayer(screenWidth, screenHeight, 1.0, resource.Provider{}.Image("images/jungle/3/4.png")),
		}), nil),
		level.NewLevel(screenWidth, screenHeight, background.NewParallax([]*background.Layer{
			background.NewLayer(screenWidth, screenHeight, 1.0, resource.Provider{}.Image("images/jungle/4/1.png")),
			background.NewLayer(screenWidth, screenHeight, 0.2, resource.Provider{}.Image("images/jungle/4/2.png")),
			background.NewLayer(screenWidth, screenHeight, 0.3, resource.Provider{}.Image("images/jungle/4/3.png")),
			background.NewLayer(screenWidth, screenHeight, 0.5, resource.Provider{}.Image("images/jungle/4/4.png")),
			background.NewLayer(screenWidth, screenHeight, 0.7, resource.Provider{}.Image("images/jungle/4/5.png")),
			background.NewLayer(screenWidth, screenHeight, 1.0, resource.Provider{}.Image("images/jungle/4/6.png")),
		}), nil),
	})
	world.stages = append(world.stages, jungleStage)

	// desert stage
	desertStage := stage.NewStage("Desert", screenWidth, screenHeight, textFaceSource, []*level.Level{
		level.NewLevel(screenWidth, screenHeight, background.NewParallax([]*background.Layer{
			background.NewLayer(screenWidth, screenHeight, 0.1, resource.Provider{}.Image("images/mountain/1/5.png")),
			background.NewLayer(screenWidth, screenHeight, 0.3, resource.Provider{}.Image("images/mountain/1/4.png")),
			background.NewLayer(screenWidth, screenHeight, 0.5, resource.Provider{}.Image("images/mountain/1/3.png")),
			background.NewLayer(screenWidth, screenHeight, 0.5, resource.Provider{}.Image("images/mountain/1/2.png")),
			background.NewLayer(screenWidth, screenHeight, 1.0, resource.Provider{}.Image("images/mountain/1/1.png")),
		}), nil),
		level.NewLevel(screenWidth, screenHeight, background.NewParallax([]*background.Layer{
			background.NewLayer(screenWidth, screenHeight, 0.1, resource.Provider{}.Image("images/mountain/2/5.png")),
			background.NewLayer(screenWidth, screenHeight, 0.3, resource.Provider{}.Image("images/mountain/2/4.png")),
			background.NewLayer(screenWidth, screenHeight, 0.5, resource.Provider{}.Image("images/mountain/2/3.png")),
			background.NewLayer(screenWidth, screenHeight, 0.5, resource.Provider{}.Image("images/mountain/2/2.png")),
			background.NewLayer(screenWidth, screenHeight, 1.0, resource.Provider{}.Image("images/mountain/2/1.png")),
		}), nil),
	})
	world.stages = append(world.stages, desertStage)
	return &world
}

func (world *World) Update() {
	stg := world.stages[world.currentStage]
	stg.Begin()
	stg.Update()
	if stg.Kind() == stage.KindGame {

	}

	// stage change over check
	if world.fading {
		world.updateFade()
		return
	}

	// game updates
	world.player.Update()
	_, lvl := stg.CurrentLevel()

	// level clear updates
	if lvl != nil && lvl.Clear() && !world.fading {
		if !world.player.WalkingToLevelExit() {
			world.fading = true
			world.fadeIn = false             // Start fade-out
			world.transitionComplete = false // Reset transitionComplete
		}
	}
}

func (world *World) Draw(screen *ebiten.Image) {
	world.stages[world.currentStage].Draw(screen)

	if world.stages[world.currentStage].Kind() != stage.KindGame {
		return
	}

	// game drawings
	world.player.Draw(screen)

	if world.fading {
		alpha := uint8(world.fadeAlpha * 255)
		vector.DrawFilledRect(screen, 0, 0, float32(float64(screen.Bounds().Dx())), float32(float64(screen.Bounds().Dy())), color.RGBA{A: alpha}, false)
	}
}

func (world *World) updateFade() {
	if world.fadeIn {
		// Fade-in phase
		world.fadeAlpha -= world.fadeSpeed
		if world.fadeAlpha <= 0 {
			world.fadeAlpha = 0
			world.fading = false            // Stop fading after fade-in is done
			world.transitionComplete = true // Transition is complete
		}
	} else {
		// Fade-out phase
		world.fadeAlpha += world.fadeSpeed
		if world.fadeAlpha >= 1 {
			world.fadeAlpha = 1
			world.fadeIn = true // Start fade-in
			if !world.transitionComplete {
				if !world.stages[world.currentStage].NextLevel() {
					if world.currentStage < len(world.stages)-1 {
						world.currentStage++
					} else {
						// game end, user must have completed the game by now
						fmt.Println("game won")
					}
				}
			}
		}
	}
}
