package fonts

import (
	"bytes"
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	//go:embed manaspace/manaspc.ttf
	ManaSpace []byte
)

const (
	DefaultTextSize = 48
)

// LoadFont loads the embedded font and returns a font face
func LoadFont(fontBytes []byte) (*text.GoTextFaceSource, error) {
	return text.NewGoTextFaceSource(bytes.NewReader(fontBytes))
}
