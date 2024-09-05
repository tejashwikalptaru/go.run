package background

import (
	"bytes"
	"image"
)

func GetImage(data []byte) (image.Image, error) {
	img, _, decodeErr := image.Decode(bytes.NewReader(data))
	if decodeErr != nil {
		return nil, decodeErr
	}
	return img, nil
}
