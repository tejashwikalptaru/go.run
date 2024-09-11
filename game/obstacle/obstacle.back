package enemy

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game/character"
	"github.com/tejashwikalptaru/go.run/resource"
)

type obstacleType string

const (
	obstacleTypeSnake    obstacleType = "snake"
	obstacleTypeHyena    obstacleType = "hyena"
	obstacleTypeScorpio  obstacleType = "scorpio"
	obstacleTypeVulture  obstacleType = "vulture"
	obstacleTypeMummy    obstacleType = "mummy"
	obstacleTypeDeceased obstacleType = "deceased"
)

const obstacleSpriteSize = 48

type obstacleSpriteInfo struct {
	frames          []*ebiten.Image
	width           float64
	height          float64
	collisionTop    float64
	collisionLeft   float64
	collisionWidth  float64
	collisionHeight float64
}

type obstacleItem struct {
	obstacleType    obstacleType
	xPosition       float64
	speed           float64
	frameIndex      int
	frameCount      int
	height          float64
	width           float64
	yPosition       float64
	passed          bool
	isPowerUpObject bool
}

type Obstacle struct {
	rng            *rand.Rand
	player         *character.Player
	obstacleImages map[obstacleType]obstacleSpriteInfo
	obstacles      []obstacleItem
	groundY        float64
	minObstacleGap float64
	maxObstacleGap float64
	screenWidth    float64
	obstacleSpeed  float64
	frameDelay     int
	scaleFactor    float64
	maxObstacles   int
	debug          bool
}

func NewObstacle(screenWidth, groundY float64, player *character.Player, rng *rand.Rand, maxObstacles int, debug bool) (*Obstacle, error) {
	obstacle := &Obstacle{
		minObstacleGap: 250,
		maxObstacleGap: 400,
		screenWidth:    screenWidth,
		rng:            rng,
		obstacleSpeed:  5,
		groundY:        groundY,
		player:         player,
		frameDelay:     5,
		scaleFactor:    1.5,
		maxObstacles:   maxObstacles,
		debug:          debug,
	}
	obstacleImagesErr := obstacle.loadObstacleSprites()
	if obstacleImagesErr != nil {
		return nil, obstacleImagesErr
	}
	obstacle.Prepare()
	return obstacle, nil
}

// loadFrames splits a sprite sheet into individual frames
func (o *Obstacle) loadFrames(img *ebiten.Image, frameWidth, frameHeight, frameCount int) []*ebiten.Image {
	frames := make([]*ebiten.Image, frameCount)
	for i := 0; i < frameCount; i++ {
		frame, ok := img.SubImage(image.Rect(i*frameWidth, 0, (i+1)*frameWidth, frameHeight)).(*ebiten.Image)
		if ok && o.debug {
			fmt.Println("failed to load sub image for obstacle")
		}
		frames[i] = frame
	}
	return frames
}

func (o *Obstacle) loadObstacleSprites() error {
	// Store frames and dimensions in the map
	o.obstacleImages = map[obstacleType]obstacleSpriteInfo{
		obstacleTypeDeceased: {
			frames:          o.loadFrames(resource.Provider{}.Image("sprites/enemy/Deceased_walk.png"), 48, 48, 6),
			width:           obstacleSpriteSize,
			height:          obstacleSpriteSize,
			collisionTop:    12,
			collisionLeft:   40,
			collisionWidth:  25,
			collisionHeight: 60,
		},
		obstacleTypeHyena: {
			frames:          o.loadFrames(resource.Provider{}.Image("sprites/enemy/Hyena_walk.png"), 48, 48, 6),
			width:           obstacleSpriteSize,
			height:          obstacleSpriteSize,
			collisionTop:    30,
			collisionLeft:   0,
			collisionWidth:  55,
			collisionHeight: 45,
		},
		obstacleTypeMummy: {
			frames:          o.loadFrames(resource.Provider{}.Image("sprites/enemy/Mummy_walk.png"), 48, 48, 6),
			width:           obstacleSpriteSize,
			height:          obstacleSpriteSize,
			collisionTop:    12,
			collisionLeft:   35,
			collisionWidth:  30,
			collisionHeight: 60,
		},
		obstacleTypeScorpio: {
			frames:          o.loadFrames(resource.Provider{}.Image("sprites/enemy/Scorpio_walk.png"), 48, 48, 4),
			width:           obstacleSpriteSize,
			height:          obstacleSpriteSize,
			collisionTop:    35,
			collisionLeft:   20,
			collisionWidth:  50,
			collisionHeight: 40,
		},
		obstacleTypeSnake: {
			frames:          o.loadFrames(resource.Provider{}.Image("sprites/enemy/Snake_walk.png"), 48, 48, 4),
			width:           obstacleSpriteSize,
			height:          obstacleSpriteSize,
			collisionTop:    45,
			collisionLeft:   20,
			collisionWidth:  50,
			collisionHeight: 40,
		},
		obstacleTypeVulture: {
			frames:          o.loadFrames(resource.Provider{}.Image("sprites/enemy/Vulture_walk.png"), 48, 48, 4),
			width:           obstacleSpriteSize,
			height:          obstacleSpriteSize,
			collisionTop:    26,
			collisionLeft:   10,
			collisionWidth:  70,
			collisionHeight: 45,
		},
	}
	return nil
}

func (o *Obstacle) randomObstacleType(rng *rand.Rand) obstacleType {
	types := []obstacleType{
		obstacleTypeSnake,
		obstacleTypeHyena,
		obstacleTypeScorpio,
		obstacleTypeVulture,
		obstacleTypeMummy,
		obstacleTypeDeceased,
	}
	return types[rng.Intn(len(types))]
}

// Prepare creates the obstacles
func (o *Obstacle) Prepare() {
	o.obstacles = []obstacleItem{} // Clear any existing obstacles

	var lastX = o.screenWidth + 300
	for i := 0; i < o.maxObstacles; i++ {
		obstacleType := o.randomObstacleType(o.rng)

		// Get the dimensions from the obstacleImages map
		spriteInfo := o.obstacleImages[obstacleType]
		obstacleWidth := spriteInfo.width * o.scaleFactor
		obstacleHeight := spriteInfo.height * o.scaleFactor

		// Calculate the Y position based on obstacle type (e.g., flying or ground-level)
		var yPosition float64
		if obstacleType == obstacleTypeVulture {
			yPosition = o.groundY - 100 - obstacleHeight // Flying obstacle above the ground
		} else {
			yPosition = o.groundY - obstacleHeight // Ground-level obstacles
		}

		// Create obstacle with random gap
		gap := o.rng.Float64()*(o.maxObstacleGap-o.minObstacleGap) + o.minObstacleGap
		lastX += gap

		newObstacle := obstacleItem{
			xPosition:    lastX,
			speed:        o.obstacleSpeed,
			obstacleType: obstacleType,
			width:        obstacleWidth,
			height:       obstacleHeight,
			yPosition:    yPosition,
		}
		o.obstacles = append(o.obstacles, newObstacle)
	}
}

// IncreaseSpeed increases the speed of the obstacles as the player progresses to new levels
func (o *Obstacle) IncreaseSpeed() {
	o.obstacleSpeed += 1.0
}

func (o *Obstacle) Reset() {
	o.obstacleSpeed = 5
	o.Prepare()
}

// filterObstacles removes obstacles that have moved off-screen
func (o *Obstacle) filterObstacles() []obstacleItem {
	var filtered []obstacleItem
	for _, obs := range o.obstacles {
		// Remove obstacles that have moved off-screen (to the left)
		if obs.xPosition > -obs.width {
			filtered = append(filtered, obs)
		}
	}
	return filtered
}

// collisionDetected checks for a collision between the player and an obstacle
func (o *Obstacle) collisionDetected(obs *obstacleItem) bool {
	// Player's collision boundaries
	playerLeft := 40 + o.player.CollisionLeft()
	playerRight := playerLeft + (o.player.CollisionWidth())
	playerTop := o.player.YPosition() + o.player.CollisionTop()
	playerBottom := playerTop + (o.player.CollisionHeight())

	// Get the obstacle's sprite information
	spriteInfo := o.obstacleImages[obs.obstacleType]

	// Fallback to full size if the collision values are 0
	collisionLeft := spriteInfo.collisionLeft
	if collisionLeft == 0 {
		collisionLeft = 0 // Default to the sprite's leftmost side
	}

	collisionTop := spriteInfo.collisionTop
	if collisionTop == 0 {
		collisionTop = 0 // Default to the sprite's topmost side
	}

	collisionWidth := spriteInfo.collisionWidth
	if collisionWidth == 0 {
		collisionWidth = spriteInfo.width // Use the full sprite width if not specified
	}

	collisionHeight := spriteInfo.collisionHeight
	if collisionHeight == 0 {
		collisionHeight = spriteInfo.height // Use the full sprite height if not specified
	}

	// Obstacle's collision boundaries (using the actual collision box or full size if not provided)
	obstacleRight := obs.xPosition + collisionLeft + collisionWidth
	obstacleLeft := obs.xPosition + collisionLeft
	obstacleTop := obs.yPosition + collisionTop
	obstacleBottom := obstacleTop + collisionHeight

	// Check for collision between the player and the obstacle's collision area
	xOverlap := playerRight > obstacleLeft && 50 < obstacleRight
	yOverlap := playerBottom > obstacleTop && playerTop < obstacleBottom

	return xOverlap && yOverlap
}

// cleared returns true when an obstacle was completely passed by player
func (o *Obstacle) cleared() bool {
	for i := range o.obstacles {
		// Check if the obstacle has completely passed the player (xPosition + width is less than player's X)
		if o.obstacles[i].xPosition+o.obstacles[i].width < 40 && !o.obstacles[i].passed {
			o.obstacles[i].passed = true // Mark the obstacle as passed
			return true
		}
	}
	return false
}

// Update handles the movement of obstacles and checks for collisions
func (o *Obstacle) Update() (collision, isPowerUpObject, cleared bool) {
	for i := range o.obstacles {
		o.obstacles[i].xPosition -= o.obstacles[i].speed // Move the obstacle to the left

		// Update frame for animation based on frame delay
		o.obstacles[i].frameCount++
		if o.obstacles[i].frameCount >= o.frameDelay {
			// Get the total number of frames for this obstacle type
			totalFrames := len(o.obstacleImages[o.obstacles[i].obstacleType].frames)
			// Cycle through the frames for animation
			o.obstacles[i].frameIndex = (o.obstacles[i].frameIndex + 1) % totalFrames
			o.obstacles[i].frameCount = 0
		}
	}
	o.obstacles = o.filterObstacles() // Remove obstacles that have moved off-screen

	// Check for collisions with each obstacle
	for _, obs := range o.obstacles {
		if o.collisionDetected(&obs) {
			// a collision is detected
			return true, obs.isPowerUpObject, false
		}
	}
	return false, false, o.cleared()
}

// Draw renders the obstacles on the screen
func (o *Obstacle) Draw(screen *ebiten.Image) {
	for _, obs := range o.obstacles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(o.scaleFactor, o.scaleFactor)
		op.GeoM.Translate(obs.xPosition, obs.yPosition)

		// Draw the current frame for the obstacle using the sprite info
		currentFrame := o.obstacleImages[obs.obstacleType].frames[obs.frameIndex]
		screen.DrawImage(currentFrame, op)

		if o.debug {
			// Visualise the collision box for debugging
			spriteInfo := o.obstacleImages[obs.obstacleType]

			// Determine the collision box
			collisionLeft := spriteInfo.collisionLeft
			if collisionLeft == 0 {
				collisionLeft = 0 // Default to the sprite's leftmost side
			}

			collisionTop := spriteInfo.collisionTop
			if collisionTop == 0 {
				collisionTop = 0 // Default to the sprite's topmost side
			}

			collisionWidth := spriteInfo.collisionWidth
			if collisionWidth == 0 {
				collisionWidth = spriteInfo.width // Use the full sprite width if not specified
			}

			collisionHeight := spriteInfo.collisionHeight
			if collisionHeight == 0 {
				collisionHeight = spriteInfo.height // Use the full sprite height if not specified
			}

			vector.DrawFilledRect(
				screen,
				float32(obs.xPosition+collisionLeft), // X position
				float32(obs.yPosition+collisionTop),  // YPosition position
				float32(collisionWidth),              // Width of the obstacle
				float32(collisionHeight),             // Height of the obstacle
				color.RGBA{R: 255, A: 128},           // Color of the rectangle (Red with 50% transparency)
				false,
			)
		}
	}
}
