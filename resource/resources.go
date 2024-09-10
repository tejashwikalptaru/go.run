package resource

import (
	"bytes"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image"
	_ "image/png"
	"io"
	"io/fs"
)

//go:embed fonts images music sprites
var FS embed.FS

type Provider struct{}

func (r Provider) Reader(path string) *bytes.Reader {
	f, err := FS.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(f fs.File) {
		closeErr := f.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}(f)
	fileContent, readErr := io.ReadAll(f)
	if readErr != nil {
		panic(readErr)
	}
	return bytes.NewReader(fileContent)
}

func (r Provider) Image(path string) *ebiten.Image {
	/****
	Tip of the day, at least for me :)
	How to find the frame size from a sprite image?
	Divide the height of the image by the number of rows. That’s the frame height.
	Divide the width of the image by the number of columns. That’s the frame width.

	So the image is 215 x 146.
	215 / 5 (number of rows) = 43
	146 / 8 (number of columns) = 18.25
	*/

	reader := r.Reader(path)
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func (r Provider) TextFaceSource(path string) *text.GoTextFaceSource {
	reader := r.Reader(path)
	tfs, err := text.NewGoTextFaceSource(reader)
	if err != nil {
		panic(err)
	}
	return tfs
}
