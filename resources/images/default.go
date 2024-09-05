package images

import (
	_ "embed"
	"image"
	"image/png"
)

var (
	//go:embed background-1.png
	BackgroundOne []byte

	//go:embed background-2.png
	BackgroundTwo []byte

	//go:embed background-3.png
	BackgroundThree []byte

	//go:embed background-4.png
	BackgroundFour []byte

	//go:embed cloud.png
	Cloud []byte
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}
