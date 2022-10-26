package generatorService

import (
	"image"

	"golang.org/x/image/draw"
)

func (g *GeneratorData) ResizeImages() {
	var minHeight int = -1
	var minWidth int = -1

	for _, img := range g.Images {
		y := img.Bounds().Dy()
		x := img.Bounds().Dx()

		if minHeight == -1 || minHeight > y {
			minHeight = y
		}
		if minWidth == -1 || minWidth > x {
			minWidth = x
		}
	}

	formulas := map[string]float64{
		"horizontal": 0.0,
		"vertical":   1.0,
	}

	orientationRatio := formulas[g.orientation]

	for i, img := range g.Images {
		x := img.Bounds().Dx()
		y := img.Bounds().Dy()

		var ratio float64 = float64(x)/float64(minWidth)*(orientationRatio) + float64(y)/float64(minHeight)*(1-orientationRatio)

		x = int(float64(x)/ratio*(1-orientationRatio) + float64(minWidth)*(orientationRatio))
		y = int(float64(y)/ratio*(orientationRatio) + float64(minHeight)*(1-orientationRatio))

		g.width = x + g.width*int(1-orientationRatio)
		g.height = y + g.height*int(orientationRatio)

		dst := image.NewRGBA(image.Rect(0, 0, x, y))
		draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

		g.Images[i] = dst
	}
}

func (g *GeneratorData) ResizeGrid() {
	var tempImages = []image.Image{}

	maxHeight := -1
	count := 0

	for _, img := range g.Images {
		tempImages = append(tempImages, img)
		y := img.Bounds().Dy()

		if maxHeight < y {
			maxHeight = y
		}

		if len(tempImages) == 2 {
			width := img.Bounds().Dx()
			for _, img := range tempImages {
				y := img.Bounds().Dy()

				dst := image.NewRGBA(image.Rect(0, 0, width, maxHeight))
				draw.NearestNeighbor.Scale(dst, dst.Rect, img, image.Rect(0, int((y-maxHeight)/2), width, maxHeight+int((y-maxHeight)/2)), draw.Over, nil)
				g.Images[count] = dst
				count++
			}

			tempImages = nil
			maxHeight = 0
		}
	}

	g.width = g.Images[0].Bounds().Dx() * 2
	g.height = g.Images[0].Bounds().Dy() * len(g.Images) / 2

}
