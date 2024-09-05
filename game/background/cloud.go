package background

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go-dino/resources/images"
	"math/rand"
)

type cloudItem struct {
	X     float64
	Y     float64
	Speed float64
}

type Cloud struct {
	cloudWidth    float64
	minCloudY     float64
	maxCloudY     float64
	minCloudSpeed float64
	maxCloudSpeed float64
	screenWidth   float64
	screenHeight  float64
	rng           *rand.Rand
	cloudImage    *ebiten.Image
	clouds        []cloudItem
}

func NewCloud(screenWidth, screenHeight, numClouds int, rng *rand.Rand) (*Cloud, error) {
	image, imageErr := GetImage(images.Cloud)
	if imageErr != nil {
		return nil, imageErr
	}
	cloud := &Cloud{
		cloudWidth:    100,
		minCloudY:     5,
		maxCloudY:     float64(screenHeight / 4),
		minCloudSpeed: 0.5,
		maxCloudSpeed: 1.0,
		screenWidth:   float64(screenWidth),
		screenHeight:  float64(screenHeight),
		rng:           rng,
		cloudImage:    ebiten.NewImageFromImage(image),
	}
	for i := 0; i < numClouds; i++ {
		cloud.addCloud()
	}
	return cloud, nil
}

// addCloud adds a new cloud at a random height and speed
func (c *Cloud) addCloud() {
	c.clouds = append(c.clouds, cloudItem{
		X:     c.screenWidth,                                                       // Start the cloud on the right side of the screen
		Y:     c.minCloudY + c.rng.Float64()*(c.maxCloudY-c.minCloudY),             // Restrict cloud Y position to upper part of screen
		Speed: c.minCloudSpeed + c.rng.Float64()*(c.maxCloudSpeed-c.minCloudSpeed), // Random speed in the range
	})
}

func (c *Cloud) Update() {
	for i := range c.clouds {
		c.clouds[i].X -= c.clouds[i].Speed // Move the cloud to the left
		if c.clouds[i].X < -c.cloudWidth { // If the cloud goes off-screen
			c.clouds[i].X = c.screenWidth                          // Reposition it to the right
			c.clouds[i].Y = c.rng.Float64() * (c.screenHeight / 2) // Random height
		}
	}
}

func (c *Cloud) Draw(screen *ebiten.Image) {
	for _, cloud := range c.clouds {
		cloudOp := &ebiten.DrawImageOptions{}
		cloudOp.GeoM.Translate(cloud.X, cloud.Y) // Translate the cloud to its position
		screen.DrawImage(c.cloudImage, cloudOp)  // Draw the cloud
	}
}
