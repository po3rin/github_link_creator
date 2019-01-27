package pipeline

import (
	"image"
	"image/draw"
	"os"

	"github.com/po3rin/github_link_creator/config"
)

func getBaseImg() (image.Image, error) {
	file, err := os.Open(config.BaseImgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// SynthesizeToBase synthesize image to base image.
func SynthesizeToBase(img image.Image) (image.Image, error) {
	baseImg, err := getBaseImg()
	if err != nil {
		return nil, err
	}

	startPointLogo := image.Point{20, 24}
	logoRectangle := image.Rectangle{startPointLogo, startPointLogo.Add(img.Bounds().Size())}
	originRectangle := image.Rectangle{image.Point{0, 0}, baseImg.Bounds().Size()}

	rgba := image.NewRGBA(originRectangle)
	draw.Draw(rgba, originRectangle, baseImg, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, logoRectangle, img, image.Point{0, 0}, draw.Over)

	return rgba, nil
}
