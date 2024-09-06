package music

import _ "embed"

var (
	//go:embed sinnesloschen-beam-117362_isak_pixabay.mp3
	Background []byte

	//go:embed cartoon-jump-6462.mp3
	Jump []byte

	//go:embed doorhit-98828.mp3
	Collision []byte
)
