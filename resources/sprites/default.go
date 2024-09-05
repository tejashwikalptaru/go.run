package sprites

import (
	_ "embed"
	"image"
	"image/png"
)

var (
	//go:embed runner.png
	Runner []byte
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}
