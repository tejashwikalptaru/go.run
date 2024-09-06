package resources

import (
	"bytes"
	"image"
)

/****
Tip of the day, at least for me :)
How to find the frame size from a sprite image?
Divide the height of the image by the number of rows. That’s the frame height.
Divide the width of the image by the number of columns. That’s the frame width.

So the image is 215 x 146.
215 / 5 (number of rows) = 43
146 / 8 (number of columns) = 18.25
*/

func GetImage(data []byte) (image.Image, error) {
	img, _, decodeErr := image.Decode(bytes.NewReader(data))
	if decodeErr != nil {
		return nil, decodeErr
	}
	return img, nil
}
