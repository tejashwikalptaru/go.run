package background

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Layer struct {
	image          *ebiten.Image
	scrollSpeed    float64 // Speed of the scrolling layer
	positionX      float64 // Position of the layer on the X-axis
	scaleX, scaleY float64 // Scaling factor
	width, height  float64 // Image height and width
}

type Parallax struct {
	layers []*Layer
}

func NewLayer(screenWidth, screenHeight, scrollSpeed float64, image *ebiten.Image) *Layer {
	width, height := float64(image.Bounds().Dx()), float64(image.Bounds().Dy())
	return &Layer{
		image:       image,
		scrollSpeed: scrollSpeed,
		width:       width,
		height:      height,
		scaleX:      screenWidth / width,
		scaleY:      screenHeight / height,
	}
}

// NewParallax creates a parallax effect with multiple layers
func NewParallax(layers []*Layer) *Parallax {
	return &Parallax{layers: layers}
}

// Update updates the position of each parallax layer
func (p *Parallax) Update() {
	for _, layer := range p.layers {
		// Move the layer to the left based on its scroll speed
		layer.positionX -= layer.scrollSpeed

		// Reset position when the scaled image has fully moved off the screen
		if layer.positionX <= -layer.width*layer.scaleX {
			layer.positionX += layer.width * layer.scaleX
		}
	}
}

func (p *Parallax) Draw(screen *ebiten.Image) {
	for _, layer := range p.layers {
		// Draw the first image
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(layer.scaleX, layer.scaleY)
		op.GeoM.Translate(layer.positionX, 0)
		screen.DrawImage(layer.image, op)

		// Draw the second image right after the first one for seamless scrolling
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Scale(layer.scaleX, layer.scaleY)
		op2.GeoM.Translate(layer.positionX+layer.width*layer.scaleX, 0)
		screen.DrawImage(layer.image, op2)
	}
}
