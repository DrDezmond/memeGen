package generatorService

import (
	"image"

	"golang.org/x/image/draw"
)

func (g *GeneratorData) CombineImages() image.Image {

	bigImage := image.Rectangle{image.Point{0, 0}, image.Point{g.width, g.height}}
	rgba := image.NewRGBA(bigImage)

	x := 0
	y := 0
	for i, img := range g.Images {

		draw.Draw(rgba, image.Rectangle{image.Point{x, y}, image.Point{img.Bounds().Dx() + x, img.Bounds().Dy() + y}}, img, image.Point{0, 0}, draw.Src)

		if g.orientation == "vertical" {
			y += img.Bounds().Dy()
		} else {
			x += img.Bounds().Dx()
			if i%2 == 1 && g.orientation == "grid" {
				y += img.Bounds().Dy()
				x = 0
			}
		}

	}

	return rgba
}
