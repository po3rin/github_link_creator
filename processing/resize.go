package processing

import (
	"image"

	"github.com/nfnt/resize"
)

// ResizeImg resize github user image.
func ResizeImg(img image.Image) image.Image {
	m := resize.Resize(248, 0, img, resize.Lanczos3)
	return m
}
