package config

import (
	"image/color"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/po3rin/github_link_creator/static"
)

// Text configs in image.
type Text struct {
	Size          float64
	Color         color.RGBA
	PosHorizontal int
	PosVertical   int
}

var (
	// Title is github repo name.
	Title Text
	// FirstDescription is github repo description in First line.
	FirstDescription Text
	// SecondDescription is github repo description in First line.
	SecondDescription Text
	// Star is github repo star.
	Star Text
	// Fork is github repo fork.
	Fork Text
	// Font is common font.
	Font *truetype.Font
)

var (
	// BaseImgPath is base image path.
	BaseImgPath = "/base.png"
)

func init() {
	setFont()

	Title.Size = 34.0
	Title.Color = color.RGBA{R: 3, G: 102, B: 214, A: 255}
	Title.PosHorizontal = 286
	Title.PosVertical = 80

	FirstDescription.Size = 24.0
	FirstDescription.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	FirstDescription.PosHorizontal = 286
	FirstDescription.PosVertical = 140

	SecondDescription.Size = 24.0
	SecondDescription.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	SecondDescription.PosHorizontal = 286
	SecondDescription.PosVertical = 180

	Star.Size = 28.0
	Star.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	Star.PosHorizontal = 332
	Star.PosVertical = 256

	Fork.Size = 28.0
	Fork.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	Fork.PosHorizontal = 456
	Fork.PosVertical = 256
}

func setFont() {
	file, err := static.Assets.Open("/mplus-1c-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	fontBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	Font = f
}
