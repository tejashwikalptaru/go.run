package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Player struct {
	height    float64
	width     float64
	y         float64
	velocityY float64
	gravity   float64
	isJumping bool
	groundY   float64
}

func NewPlayer(groundY float64) *Player {
	height := 60.0
	return &Player{
		height:    height,
		width:     40,
		velocityY: 0,
		gravity:   0.6,
		groundY:   groundY,
		y:         groundY - height,
	}
}

func (p *Player) Width() float64 {
	return p.width
}

func (p *Player) Height() float64 {
	return p.height
}

func (p *Player) Y() float64 {
	return p.y
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !p.isJumping {
		p.velocityY = -12
		p.isJumping = true
	}

	// Apply gravity and update dino's position
	if p.isJumping {
		p.y += p.velocityY
		p.velocityY += p.gravity

		// Stop falling when reaching the ground
		if p.y >= p.groundY-p.height {
			p.y = p.groundY - p.height
			p.isJumping = false
			p.velocityY = 0
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 40, float32(p.y), float32(p.width), float32(p.height), color.RGBA{G: 128, A: 255}, false)
}
