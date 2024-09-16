package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/entity"
	"github.com/tejashwikalptaru/go.run/game/music"
	"github.com/tejashwikalptaru/go.run/resource"
)

type Player struct {
	entity.BaseEntity
	groundY          float64
	isJumping        bool
	velocityY        float64
	velocityX        float64
	runningSpeed     float64 // Running speed for horizontal movement
	maxXVelocity     float64
	walkingToExit    bool
	screenWidth      float64
	xPositionDesired float64
	gravity          float64

	jumpMusic *music.Manager

	heart entity.BaseEntity
	life  int
}

func New(screenWidth, groundY float64) *Player {
	playerEntity := entity.New(
		resource.Provider{}.Image("sprites/player.png"),
		32,
		32,
		8,
		5,
		1,
		"player",
		2,
	)

	heartEntity := entity.New(
		resource.Provider{}.Image("sprites/heart.png"),
		36,
		36,
		8,
		5,
		0,
		"life",
		1,
	)

	player := Player{
		BaseEntity:       playerEntity,
		groundY:          groundY,
		screenWidth:      screenWidth,
		xPositionDesired: 40,
		gravity:          0.6,
		heart:            heartEntity,
		life:             3,
		maxXVelocity:     2.0,
		runningSpeed:     2.0,
		jumpMusic:        music.NewMusic(resource.Provider{}.Reader("music/jump.mp3")),
	}
	player.Reset()
	return &player
}

func (p *Player) Update() {
	screenLimitLeft := 0.0            // Left edge of the screen
	screenLimitRight := p.screenWidth // Right edge of the screen

	// Jump logic
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !p.isJumping && !p.walkingToExit {
		p.velocityY = -12
		p.isJumping = true
		p.jumpMusic.Play()
	}

	// Walk the player in (when not jumping or exiting the level)
	if p.XPosition() < p.xPositionDesired && !p.isJumping && !p.walkingToExit {
		p.SetXPosition(p.XPosition() + 1)
	}

	// Handle running left or right (whether jumping or on the ground, but not when walking to exit)
	if !p.walkingToExit {
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.XPosition() > screenLimitLeft {
			p.velocityX = -p.runningSpeed // Move left
		} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.XPosition() < screenLimitRight-p.Width() {
			p.velocityX = p.runningSpeed // Move right
		} else {
			p.velocityX = 0 // Stop running if no keys pressed
		}
	}

	// Apply gravity and update player's Y position when jumping
	if p.isJumping {
		p.SetYPosition(p.YPosition() + p.velocityY)
		p.velocityY += p.gravity

		// Stop falling when reaching the ground
		if p.YPosition() >= p.groundY-p.Height() {
			p.SetYPosition(p.groundY - p.Height())
			p.isJumping = false
			p.velocityY = 0
		}
	}

	// Update the player's X position for both ground running and jumping
	p.SetXPosition(p.XPosition() + p.velocityX)

	// Ensure the player does not move beyond screen boundaries (unless walking off the screen)
	if !p.walkingToExit {
		if p.XPosition() <= screenLimitLeft {
			p.SetXPosition(screenLimitLeft) // Stop at the left edge
		} else if p.XPosition() >= screenLimitRight-p.Width() {
			p.SetXPosition(screenLimitRight - p.Width()) // Stop at the right edge
		}
	}

	// Update animations and heart entity
	p.BaseEntity.Update()
	p.heart.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.BaseEntity.Draw(screen)

	for i := 0; i < p.life; i++ {
		p.heart.SetXPosition(5 + float64(i)*p.heart.Width())
		p.heart.SetYPosition(5)
		p.heart.Draw(screen)
	}
}

// Reset function brings the player back to the ground and resets jumping state.
func (p *Player) Reset() {
	p.SetXPosition(-70)
	p.SetYPosition(p.groundY - p.Height())
	p.isJumping = false
	p.velocityY = 0
}

func (p *Player) WalkingToLevelExit() bool {
	if !p.walkingToExit {
		p.walkingToExit = true
	}

	// If the player is completely off the screen, reset and finish the exit walk
	if p.XPosition() > p.screenWidth+p.Width() {
		p.walkingToExit = false
		p.Reset()
		return false // Return false to indicate the walk is completed
	}

	// Continue walking to the right by increasing X position
	p.SetXPosition(p.XPosition() + 2)

	// Return true to indicate the player is still walking to the exit
	return true
}

func (p *Player) Hurt() (died bool) {
	p.life--
	died = p.life <= 0
	return died
}

func (p *Player) Die() {
	p.jumpMusic.Stop()
}
