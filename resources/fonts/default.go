package fonts

import (
	_ "embed"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	//go:embed manaspace/manaspc.ttf
	ManaSpace []byte
)

// LoadFont loads the embedded font and returns a font face
func LoadFont(fontBytes []byte) font.Face {
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	const dpi = 72
	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}

	return fontFace
}
