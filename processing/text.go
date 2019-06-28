package processing

import (
	"image"

	"github.com/po3rin/github_link_creator/config"
	"github.com/po3rin/txt2img"
)

// DrawText shorthand to draw text.
func DrawText(img image.Image, c config.Text, text string) (image.Image, error) {
	d, err := txt2img.NewDrawer(
		txt2img.Params{
			Img:               img,
			FontSize:          c.Size,
			Font:              config.Font,
			TextColor:         c.Color,
			TextPosHorizontal: c.PosHorizontal,
			TextPosVertical:   c.PosVertical,
		},
	)
	if err != nil {
		return nil, err
	}
	result, err := d.Draw(text)
	if err != nil {
		return nil, err
	}
	return result, nil
}
