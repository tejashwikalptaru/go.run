package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/entity"
	"github.com/tejashwikalptaru/go.run/resource"
)

type Player struct {
	entity.BaseEntity
	groundY          float64
	isJumping        bool
	velocityY        float64
	walkingToExit    bool
	screenWidth      float64
	xPositionDesired float64
	gravity          float64
}

func New(screenWidth, groundY float64) *Player {
	img := resource.Provider{}.Image("sprites/player.png")
	playerEntity := entity.New(img, 32, 32, 8, 5, 1, "player", 2.0)
	player := Player{BaseEntity: playerEntity, groundY: groundY, screenWidth: screenWidth, xPositionDesired: 40, gravity: 0.6}
	player.Reset()
	return &player
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !p.isJumping {
		p.velocityY = -12
		p.isJumping = true
		// play jump music here
	}

	// walk the player in
	if p.XPosition() < p.xPositionDesired {
		p.SetXPosition(p.XPosition() + 1)
	}

	// Apply gravity and update player's position
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
	p.BaseEntity.Update()
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

	if p.XPosition()+p.Width() >= p.screenWidth+100 {
		p.walkingToExit = false
		p.Reset()
		return false // Return false to indicate the walk is completed
	}

	// Continue walking to the right by increasing X position
	p.SetXPosition(p.XPosition() + 2)

	// Return true to indicate the player is still walking to the exit
	return true
}
