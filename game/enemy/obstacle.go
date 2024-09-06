package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tejashwikalptaru/go.run/game/character"
	"github.com/tejashwikalptaru/go.run/game/stage"
	"github.com/tejashwikalptaru/go.run/resources"
	"github.com/tejashwikalptaru/go.run/resources/sprites"
	"image"
	"image/color"
	"math/rand"
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
	xPosition    float64
	speed        float64
	passed       bool
	obstacleType obstacleType
	frameIndex   int
	frameCount   int
	height       float64
	width        float64
	yPosition    float64
}

type Obstacle struct {
	groundY        float64
	minObstacleGap float64
	maxObstacleGap float64
	screenWidth    float64
	obstacleSpeed  float64
	rng            *rand.Rand
	player         *character.Player
	obstacles      []obstacleItem
	obstacleImages map[obstacleType]obstacleSpriteInfo
	frameDelay     int
	scaleFactor    float64
}

func NewObstacle(screenWidth, groundY float64, player *character.Player, rng *rand.Rand) (*Obstacle, error) {
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
	}
	obstacleImagesErr := obstacle.loadObstacleSprites()
	if obstacleImagesErr != nil {
		return nil, obstacleImagesErr
	}
	// Start the first obstacle farther away to avoid instant collision
	obstacle.obstacles = []obstacleItem{}
	return obstacle, nil
}

// loadFrames splits a sprite sheet into individual frames
func (o *Obstacle) loadFrames(img *ebiten.Image, frameWidth, frameHeight, frameCount int) []*ebiten.Image {
	frames := make([]*ebiten.Image, frameCount)
	for i := 0; i < frameCount; i++ {
		frame := img.SubImage(image.Rect(i*frameWidth, 0, (i+1)*frameWidth, frameHeight)).(*ebiten.Image)
		frames[i] = frame
	}
	return frames
}

func (o *Obstacle) loadObstacleSprites() error {
	deceased, deceasedErr := resources.GetImage(sprites.DeceasedWalk)
	if deceasedErr != nil {
		return deceasedErr
	}
	hyena, hyenaErr := resources.GetImage(sprites.HyenaWalk)
	if hyenaErr != nil {
		return hyenaErr
	}
	mummy, mummyErr := resources.GetImage(sprites.MummyWalk)
	if mummyErr != nil {
		return mummyErr
	}
	scorpio, scorpioErr := resources.GetImage(sprites.ScorpioWalk)
	if scorpioErr != nil {
		return scorpioErr
	}
	snake, snakeErr := resources.GetImage(sprites.SnakeWalk)
	if snakeErr != nil {
		return snakeErr
	}
	vulture, vultureErr := resources.GetImage(sprites.VultureWalk)
	if vultureErr != nil {
		return vultureErr
	}
	// Store frames and dimensions in the map
	o.obstacleImages = map[obstacleType]obstacleSpriteInfo{
		obstacleTypeDeceased: {
			frames:          o.loadFrames(ebiten.NewImageFromImage(deceased), 48, 48, 6),
			width:           48,
			height:          48,
			collisionTop:    12,
			collisionLeft:   40,
			collisionWidth:  25,
			collisionHeight: 60,
		},
		obstacleTypeHyena: {
			frames:          o.loadFrames(ebiten.NewImageFromImage(hyena), 48, 48, 6),
			width:           48,
			height:          48,
			collisionTop:    30,
			collisionLeft:   0,
			collisionWidth:  55,
			collisionHeight: 45,
		},
		obstacleTypeMummy: {
			frames:          o.loadFrames(ebiten.NewImageFromImage(mummy), 48, 48, 6),
			width:           48,
			height:          48,
			collisionTop:    12,
			collisionLeft:   35,
			collisionWidth:  30,
			collisionHeight: 60,
		},
		obstacleTypeScorpio: {
			frames:          o.loadFrames(ebiten.NewImageFromImage(scorpio), 48, 48, 4),
			width:           48,
			height:          48,
			collisionTop:    35,
			collisionLeft:   20,
			collisionWidth:  55,
			collisionHeight: 40,
		},
		obstacleTypeSnake: {
			frames:          o.loadFrames(ebiten.NewImageFromImage(snake), 48, 48, 4),
			width:           48,
			height:          48,
			collisionTop:    45,
			collisionLeft:   20,
			collisionWidth:  55,
			collisionHeight: 40,
		},
		obstacleTypeVulture: {
			frames:          o.loadFrames(ebiten.NewImageFromImage(vulture), 48, 48, 4),
			width:           48,
			height:          48,
			collisionTop:    26,
			collisionLeft:   10,
			collisionWidth:  70,
			collisionHeight: 45,
		},
	}
	return nil
}

func (o *Obstacle) randomObstacleType(rng *rand.Rand) obstacleType {
	//types := []obstacleType{
	//	obstacleTypeSnake,
	//	obstacleTypeHyena,
	//	obstacleTypeScorpio,
	//	obstacleTypeVulture,
	//	obstacleTypeMummy,
	//	obstacleTypeDeceased,
	//}
	//return types[rng.Intn(len(types))]
	return obstacleTypeHyena
}

// IncreaseSpeed increases the speed of the obstacles as the player progresses to new levels
func (o *Obstacle) IncreaseSpeed() {
	o.obstacleSpeed += 1.0
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

// addObstacleWithGap add new obstacle with specific height, width, and YPosition position based on type
func (o *Obstacle) addObstacleWithGap() {
	lastObstacle := o.obstacles[len(o.obstacles)-1]

	// Create a new obstacle with a random type and random gap
	gap := o.rng.Float64()*(400-250) + 250
	obstacleType := o.randomObstacleType(o.rng)

	// Get the dimensions from the obstacleImages map
	spriteInfo := o.obstacleImages[obstacleType]
	obstacleWidth := spriteInfo.width * o.scaleFactor
	obstacleHeight := spriteInfo.height * o.scaleFactor

	// Calculate the YPosition position based on obstacle type (e.g., flying or ground-level)
	var yPosition float64
	if obstacleType == obstacleTypeVulture {
		// Adjust yPosition for flying obstacles
		yPosition = o.groundY - 100 - obstacleHeight // Flying obstacle above the ground, 100
	} else {
		// Ground-level obstacles, adjust to ensure they are on the ground
		yPosition = o.groundY - obstacleHeight
	}

	newObstacle := obstacleItem{
		xPosition:    lastObstacle.xPosition + gap,
		speed:        o.obstacleSpeed,
		obstacleType: obstacleType,
		width:        obstacleWidth,  // Use scaled width
		height:       obstacleHeight, // Use scaled height
		yPosition:    yPosition,      // Use adjusted YPosition position
	}
	o.obstacles = append(o.obstacles, newObstacle)
}

// collisionDetected checks for a collision between the player and an obstacle
func (o *Obstacle) collisionDetected(obs obstacleItem) bool {
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

// TrackJumpsAndScore checks if the player has successfully jumped over an obstacle
func (o *Obstacle) TrackJumpsAndScore(level *stage.Level) bool {
	for i := range o.obstacles {
		// Check if the obstacle has completely passed the player (xPosition + width is less than player's X)
		if o.obstacles[i].xPosition+o.obstacles[i].width < 40 && !o.obstacles[i].passed {
			o.obstacles[i].passed = true // Mark the obstacle as passed
			level.IncreaseJump()         // Increment jump count
			level.IncreaseScore()        // Increment the score by 10

			// Check for level progression
			return true
		}
	}
	return false
}

// Reset clears the current obstacles and resets the obstacle list
func (o *Obstacle) Reset() {
	// Choose a random obstacle type for the first obstacle
	obstacleType := o.randomObstacleType(o.rng)

	// Get the dimensions from the obstacleImages map
	spriteInfo := o.obstacleImages[obstacleType]
	obstacleWidth := spriteInfo.width * o.scaleFactor
	obstacleHeight := spriteInfo.height * o.scaleFactor

	// Calculate the YPosition position based on obstacle type (e.g., flying or ground-level)
	var yPosition float64
	if obstacleType == obstacleTypeVulture {
		// Adjust yPosition for flying obstacles
		yPosition = o.groundY - 100 - obstacleHeight // Flying obstacle above the ground
	} else {
		// Ground-level obstacles, adjust to ensure they are on the ground
		yPosition = o.groundY - obstacleHeight
	}

	// Start with one obstacle far away
	o.obstacles = []obstacleItem{
		{
			xPosition:    o.screenWidth + 300,
			speed:        o.obstacleSpeed,
			obstacleType: obstacleType,
			width:        obstacleWidth,
			height:       obstacleHeight,
			yPosition:    yPosition,
		},
	}
}

// Update handles the movement of obstacles and checks for collisions
func (o *Obstacle) Update() bool {
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

	// Ensure proper gap between consecutive obstacles
	if len(o.obstacles) == 0 || (o.obstacles[len(o.obstacles)-1].xPosition < o.screenWidth-o.minObstacleGap) {
		o.addObstacleWithGap() // Add new obstacle with appropriate gap
	}

	// Check for collisions with each obstacle
	for _, obs := range o.obstacles {
		if o.collisionDetected(obs) {
			return true // Trigger game over if a collision is detected
		}
	}
	return false
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
