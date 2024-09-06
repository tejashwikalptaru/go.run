package sprites

import (
	_ "embed"
	"image"
	"image/png"
)

var (
	//go:embed runner.png
	Runner []byte

	//go:embed enemy/Snake_walk.png
	SnakeWalk []byte

	//go:embed enemy/Hyena_walk.png
	HyenaWalk []byte

	//go:embed enemy/Mummy_walk.png
	MummyWalk []byte

	//go:embed enemy/Scorpio_walk.png
	ScorpioWalk []byte

	//go:embed enemy/Vulture_walk.png
	VultureWalk []byte

	//go:embed enemy/Deceased_walk.png
	DeceasedWalk []byte
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}
