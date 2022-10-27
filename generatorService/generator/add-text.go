package generatorService

import (
	"image"
	"strings"

	"github.com/golang/freetype"
)

func (g *GeneratorData) AddText() {
	face, f, _ := LoadFont()
	c := freetype.NewContext()
	c.SetFont(f)
	c.SetFontSize(g.fontSize)
	c.SetSrc(image.White)

	for i, img := range g.Images {
		rgba := img.(*image.RGBA)
		c.SetDst(rgba)
		c.SetClip(rgba.Bounds())

		for j, val := range (g.texts)[i] {
			res := strings.Split(val, "\n")
			rows := len(res)

			for count, text_string := range res {
				c.SetSrc(image.Black)
				c.SetFontSize(g.fontSize)

				for deltaX := -1; deltaX < 2; deltaX++ {
					for deltaY := -1; deltaY < 2; deltaY++ {
						outlineX := (rgba.Bounds().Dx())/2 + deltaX
						outlineY := j*(rgba.Bounds().Dy()-int(g.fontSize)*rows-10) + int(g.fontSize)*(count+1) + deltaY

						DrawText(c, text_string, outlineX, outlineY, face)
					}
				}

				c.SetSrc(image.White)
				c.SetFontSize(g.fontSize)
				imageX := (rgba.Bounds().Dx()) / 2
				imageY := j*(rgba.Bounds().Dy()-int(g.fontSize)*rows-10) + int(g.fontSize)*(count+1)

				DrawText(c, text_string, imageX, imageY, face)
			}
		}
	}
}
