package character

import (
	"bytes"
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/music"
	"github.com/tejashwikalptaru/go.run/resources/sprites"
)

type Player struct {
	sprite           *ebiten.Image
	musicManager     *music.Manager
	xPosition        float64
	frameDelay       int
	collisionWidth   float64
	gravity          float64
	groundY          float64
	scaleFactor      float64
	collisionTop     float64
	collisionLeft    float64
	width            float64
	frameOY          int
	velocityY        float64
	xPositionDesired float64
	yPosition        float64
	frameWidth       int
	frameHeight      int
	frameIndex       int
	collisionHeight  float64
	frameCount       int
	height           float64
	screenWidth      float64
	walkingToExit    bool
	isJumping        bool
	debug            bool
}

func NewPlayer(screenWidth, groundY float64, musicManager *music.Manager, debug bool) (*Player, error) {
	height := 64.0
	// Load the player sprite sheet (runner animation)
	img, _, err := image.Decode(bytes.NewReader(sprites.Runner))
	if err != nil {
		return nil, err
	}
	return &Player{
		height:           height,
		width:            64,
		velocityY:        0,
		gravity:          0.6,
		groundY:          groundY,
		yPosition:        groundY - height,
		xPosition:        -70,
		xPositionDesired: 40,
		frameOY:          32,
		frameWidth:       32, // Width of a single frame in the sprite sheet
		frameHeight:      32, // Height of a single frame in the sprite sheet
		frameIndex:       0,
		frameDelay:       5, // Animation speed
		frameCount:       0,
		sprite:           ebiten.NewImageFromImage(img),
		scaleFactor:      2.0,
		collisionTop:     10,
		collisionLeft:    20,
		collisionWidth:   25,
		collisionHeight:  55,
		walkingToExit:    false,
		musicManager:     musicManager,
		screenWidth:      screenWidth,
		debug:            debug,
	}, nil
}

func (p *Player) Width() float64 {
	return p.width
}

func (p *Player) Height() float64 {
	return p.height
}

func (p *Player) YPosition() float64 {
	return p.yPosition
}

func (p *Player) CollisionTop() float64 {
	return p.collisionTop
}

func (p *Player) CollisionLeft() float64 {
	return p.collisionLeft
}

func (p *Player) CollisionWidth() float64 {
	return p.collisionWidth
}

func (p *Player) CollisionHeight() float64 {
	return p.collisionHeight
}

func (p *Player) ScaleFactor() float64 {
	return p.scaleFactor
}

func (p *Player) IsImmune() bool {
	return false
}

// Reset function brings the player back to the ground and resets jumping state.
func (p *Player) Reset() {
	p.xPosition = -70
	p.yPosition = p.groundY - p.height
	p.isJumping = false
	p.velocityY = 0
	p.frameIndex = 0
	p.frameCount = 0
}

func (p *Player) WalkingToLevelExit() bool {
	// If walking to exit hasn't started yet, start walking
	if !p.walkingToExit {
		p.walkingToExit = true
	}

	// Check if player has reached the right edge of the screen
	if p.xPosition+p.width >= p.screenWidth+100 {
		// Player reached the end, stop walking
		p.walkingToExit = false
		return false // Return false to indicate the walk is completed
	}

	// Continue walking to the right by increasing X position
	// Increment by a certain speed to simulate walking
	walkingSpeed := 2.0
	p.xPosition += walkingSpeed

	// Return true to indicate the player is still walking to the exit
	return true
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !p.isJumping {
		p.velocityY = -12
		p.isJumping = true
		p.musicManager.PlayJumpSound()
	}

	// walk the player in
	if p.xPosition < p.xPositionDesired {
		p.xPosition++
	}

	// Apply gravity and update player's position
	if p.isJumping {
		p.yPosition += p.velocityY
		p.velocityY += p.gravity

		// Stop falling when reaching the ground
		if p.yPosition >= p.groundY-p.height {
			p.yPosition = p.groundY - p.height
			p.isJumping = false
			p.velocityY = 0
		}
	}

	// Update animation frame (cycle through the runner frames)
	p.frameCount++
	if p.frameCount >= p.frameDelay {
		p.frameIndex++
		p.frameCount = 0       // ResetToFirst frame count after updating the frame
		if p.frameIndex >= 8 { // Assuming 8 frames in the sprite sheet
			p.frameIndex = 0
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Calculate the frame position on the sprite sheet
	sx := p.frameIndex * p.frameWidth

	// Define the part of the sprite sheet to draw (one frame)
	subImage, ok := p.sprite.SubImage(image.Rect(sx, p.frameOY, sx+p.frameWidth, p.frameOY+p.frameHeight)).(*ebiten.Image)
	if ok && p.debug {
		fmt.Println("failed to load sub image for player")
	}

	// Create image drawing options
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.scaleFactor, p.scaleFactor) // Scale the sprite to make it larger
	op.GeoM.Translate(p.xPosition, p.yPosition) // Position the sprite at the player's position
	// Draw the sprite using the current frame
	screen.DrawImage(subImage, op)

	if p.debug {
		// Visualise the collision box for debugging
		collisionTop := p.collisionTop
		if collisionTop == 0 {
			collisionTop = 0
		}

		collisionLeft := p.collisionLeft
		if collisionLeft == 0 {
			collisionLeft = 0
		}

		collisionWidth := p.collisionWidth
		if collisionWidth == 0 {
			collisionWidth = p.width
		}

		collisionHeight := p.collisionHeight
		if collisionHeight == 0 {
			collisionHeight = p.height
		}

		// Draw the player's collision rectangle
		vector.DrawFilledRect(
			screen,
			float32(p.xPosition+collisionLeft), // X position with collision offset
			float32(p.yPosition+collisionTop),  // Y position with collision offset
			float32(collisionWidth),            // Scaled collision width
			float32(collisionHeight),           // Scaled collision height
			color.RGBA{R: 255, A: 128},         // Color of the rectangle (Red with 50% transparency)
			false,
		)
	}
}
