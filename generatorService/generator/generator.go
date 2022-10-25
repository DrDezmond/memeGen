package generatorService

import (
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/golang/freetype"
	"golang.org/x/image/draw"
)

type GeneratorData struct {
	imagesAddresses *[]string
	texts           *map[int][]string
	images          []image.Image
	orientation     *string
	width           int
	height          int
}

func (g *GeneratorData) InitGeneratorValues(addresses *[]string, texts *map[int][]string, orientation *string) {

	g.imagesAddresses = addresses
	g.texts = texts
	g.images = OpenImages(*g.imagesAddresses)

	orientations := [4]string{"horizontal", "vertical", "grid"}

	for _, o := range orientations {
		if o == *orientation {
			g.orientation = orientation
		}
	}
}

func (g *GeneratorData) GenerateImages() {

	if *g.orientation == "grid" {
		*g.orientation = "vertical"
		g.ResizeImages()
		*g.orientation = "grid"
		g.ResizeGrid()
	} else {
		g.ResizeImages()
	}

	g.AddText()
	resultImage := g.CombineImages()
	GenerateOutput(resultImage)
}

func (g *GeneratorData) AddText() {
	_, f, _ := LoadFont()
	c := freetype.NewContext()
	fontSize := 24.0
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetSrc(image.White)

	for i, img := range g.images {
		rgba := img.(*image.RGBA)

		c.SetDst(rgba)
		c.SetClip(rgba.Bounds())

		for j, val := range (*g.texts)[i] {
			c.SetSrc(image.Black)
			c.SetFontSize(25.0)
			outlineX := (rgba.Bounds().Dx()-len(val)*10)/2 - 1
			outlineY := j*(rgba.Bounds().Dy()-int(fontSize*1.5)) + int(fontSize) + 1

			DrawText(c, val, outlineX, outlineY)

			c.SetSrc(image.White)
			c.SetFontSize(fontSize)
			imageX := (rgba.Bounds().Dx() - len(val)*10) / 2
			imageY := j*(rgba.Bounds().Dy()-int(fontSize*1.5)) + int(fontSize)

			DrawText(c, val, imageX, imageY)

		}
	}
}

func (g *GeneratorData) ResizeImages() {
	var minHeight int = -1
	var minWidth int = -1

	for _, img := range g.images {
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

	orientationRatio := formulas[*g.orientation]

	for i, img := range g.images {
		x := img.Bounds().Dx()
		y := img.Bounds().Dy()

		var ratio float64 = float64(x)/float64(minWidth)*(orientationRatio) + float64(y)/float64(minHeight)*(1-orientationRatio)

		x = int(float64(x)/ratio*(1-orientationRatio) + float64(minWidth)*(orientationRatio))
		y = int(float64(y)/ratio*(orientationRatio) + float64(minHeight)*(1-orientationRatio))

		g.width = x + g.width*int(1-orientationRatio)
		g.height = y + g.height*int(orientationRatio)

		dst := image.NewRGBA(image.Rect(0, 0, x, y))
		draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

		g.images[i] = dst
	}
}

func (g *GeneratorData) ResizeGrid() {
	var tempImages = []image.Image{}

	maxHeight := -1
	count := 0

	for _, img := range g.images {
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
				g.images[count] = dst
				count++
			}

			tempImages = nil
			maxHeight = 0
		}
	}

	g.width = g.images[0].Bounds().Dx() * 2
	g.height = g.images[0].Bounds().Dy() * len(g.images) / 2

}

func (g *GeneratorData) CombineImages() image.Image {

	bigImage := image.Rectangle{image.Point{0, 0}, image.Point{g.width, g.height}}
	rgba := image.NewRGBA(bigImage)

	x := 0
	y := 0
	for i, img := range g.images {

		draw.Draw(rgba, image.Rectangle{image.Point{x, y}, image.Point{img.Bounds().Dx() + x, img.Bounds().Dy() + y}}, img, image.Point{0, 0}, draw.Src)

		if *g.orientation == "vertical" {
			y += img.Bounds().Dy()
		} else {
			x += img.Bounds().Dx()
			if i%2 == 1 && *g.orientation == "grid" {
				y += img.Bounds().Dy()
				x = 0
			}
		}

	}

	return rgba
}