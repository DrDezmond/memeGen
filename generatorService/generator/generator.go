package generatorService

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
)

type GeneratorData struct {
	orientation string
	texts       map[int][]string
	fontSize    float64
	Images      []image.Image
	width       int
	height      int
}

func New() *GeneratorData {
	return &GeneratorData{}
}

func (g *GeneratorData) InitGeneratorValues(texts map[int][]string, orientation string, fontSize float64) {
	g.texts = texts

	if fontSize == 0.0 {
		g.fontSize = 24.0
	} else {
		g.fontSize = fontSize
	}

	orientations := [4]string{"horizontal", "vertical", "grid"}

	for _, o := range orientations {
		if o == orientation {
			g.orientation = orientation
		}
	}
}

func (g *GeneratorData) GenerateImages() image.Image {
	g.width = 0
	g.height = 0
	if g.orientation == "grid" {
		g.orientation = "vertical"
		g.ResizeImages()
		g.orientation = "grid"
		g.ResizeGrid()
	} else {
		g.ResizeImages()
	}

	g.AddText()
	resultImage := g.CombineImages()

	return resultImage
}
