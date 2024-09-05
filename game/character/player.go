package character

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go-dino/resources/sprites"
	"image"
)

type Player struct {
	height      float64
	width       float64
	y           float64
	velocityY   float64
	gravity     float64
	isJumping   bool
	groundY     float64
	frameOY     int
	frameWidth  int
	frameHeight int
	frameIndex  int
	frameDelay  int
	frameCount  int
	sprite      *ebiten.Image
	scaleFactor float64
}

func NewPlayer(groundY float64) (*Player, error) {
	height := 64.0
	// Load the player sprite sheet (runner animation)
	img, _, err := image.Decode(bytes.NewReader(sprites.Runner))
	if err != nil {
		return nil, err
	}
	return &Player{
		height:      height,
		width:       64,
		velocityY:   0,
		gravity:     0.6,
		groundY:     groundY,
		y:           groundY - height,
		frameOY:     32,
		frameWidth:  32, // Width of a single frame in the sprite sheet
		frameHeight: 32, // Height of a single frame in the sprite sheet
		frameIndex:  0,
		frameDelay:  5, // Animation speed
		frameCount:  0,
		sprite:      ebiten.NewImageFromImage(img),
		scaleFactor: 2.0,
	}, nil
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

	// Update animation frame (cycle through the runner frames)
	p.frameCount++
	if p.frameCount >= p.frameDelay {
		p.frameIndex++
		p.frameCount = 0       // Reset frame count after updating the frame
		if p.frameIndex >= 8 { // Assuming 8 frames in the sprite sheet
			p.frameIndex = 0
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Calculate the frame position on the sprite sheet
	sx := p.frameIndex * p.frameWidth

	// Define the part of the sprite sheet to draw (one frame)
	subImage := p.sprite.SubImage(image.Rect(sx, p.frameOY, sx+p.frameWidth, p.frameOY+p.frameHeight)).(*ebiten.Image)

	// Create image drawing options
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.scaleFactor, p.scaleFactor) // Scale the sprite to make it larger
	op.GeoM.Translate(40, p.y)                  // Position the sprite at the player's position

	// Draw the sprite using the current frame
	screen.DrawImage(subImage, op)
}
