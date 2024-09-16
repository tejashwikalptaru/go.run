package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Frame struct {
	Image            *ebiten.Image
	TightBoundingBox *image.Rectangle
	dataComputed     bool
}

func (f *Frame) computeTightBoundingBox() {
	if f.dataComputed {
		return
	}
	f.dataComputed = true
	bounds := f.Image.Bounds()
	minX, minY := bounds.Max.X, bounds.Max.Y
	maxX, maxY := bounds.Min.X, bounds.Min.Y

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := f.Image.At(x, y).RGBA()
			if a > 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// If no non-transparent pixels are found, return an empty rectangle
	if minX > maxX || minY > maxY {
		rect := image.Rect(0, 0, 0, 0)
		f.TightBoundingBox = &rect
		return
	}

	// Adjust coordinates to be relative to the sub-image's bounds
	adjustedMinX := minX - bounds.Min.X
	adjustedMinY := minY - bounds.Min.Y
	adjustedMaxX := maxX - bounds.Min.X
	adjustedMaxY := maxY - bounds.Min.Y

	rect := image.Rect(adjustedMinX, adjustedMinY, adjustedMaxX+1, adjustedMaxY+1)
	f.TightBoundingBox = &rect
}
