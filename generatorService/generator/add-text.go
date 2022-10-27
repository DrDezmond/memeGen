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
			c.SetFontSize(g.fontSize)
			for deltaX := -1; deltaX < 2; deltaX++ {
				for deltaY := -1; deltaY < 2; deltaY++ {
					outlineX := (rgba.Bounds().Dx()-len(val)*int(g.fontSize/1.4))/2 + deltaX
					outlineY := j*(rgba.Bounds().Dy()-int(g.fontSize*1.5)) + int(g.fontSize) + deltaY
					DrawText(c, val, outlineX, outlineY)
				}
			}
			c.SetSrc(image.White)
			c.SetFontSize(g.fontSize)
			imageX := (rgba.Bounds().Dx() - len(val)*int(g.fontSize/1.4)) / 2
			imageY := j*(rgba.Bounds().Dy()-int(g.fontSize*1.5)) + int(g.fontSize)

			DrawText(c, val, imageX, imageY)

		}
	}
}
