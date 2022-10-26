package generatorService

import (
	"image"

	"github.com/golang/freetype"
)

func (g *GeneratorData) AddText() {
	_, f, _ := LoadFont()
	c := freetype.NewContext()
	c.SetFont(f)
	c.SetFontSize(g.fontSize)
	c.SetSrc(image.White)

	for i, img := range g.Images {
		rgba := img.(*image.RGBA)

		c.SetDst(rgba)
		c.SetClip(rgba.Bounds())

		for j, val := range (g.texts)[i] {
			c.SetSrc(image.Black)
			c.SetFontSize(25.0)
			outlineX := (rgba.Bounds().Dx()-len(val)*10)/2 - 1
			outlineY := j*(rgba.Bounds().Dy()-int(g.fontSize*1.5)) + int(g.fontSize) + 1

			DrawText(c, val, outlineX, outlineY)

			c.SetSrc(image.White)
			c.SetFontSize(g.fontSize)
			imageX := (rgba.Bounds().Dx() - len(val)*10) / 2
			imageY := j*(rgba.Bounds().Dy()-int(g.fontSize*1.5)) + int(g.fontSize)

			DrawText(c, val, imageX, imageY)

		}
	}
}
