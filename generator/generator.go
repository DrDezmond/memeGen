package generator

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"example.com/generator/generator/helpers"
	"github.com/golang/freetype"
	"golang.org/x/image/draw"
)

type GeneratorInputData struct {
	imagesAddresses *[]string
	texts           *map[int][]string
	images          []image.Image
	orientation     *string
}

func (g *GeneratorInputData) InitGeneratorValues(addresses *[]string, texts *map[int][]string, orientation *string) {

	g.imagesAddresses = addresses
	g.texts = texts

	orientations := [4]string{"single", "horizontal", "vertical", "grid"}

	for _, o := range orientations {
		if o == *orientation {
			g.orientation = orientation
		}
	}
}

func (g *GeneratorInputData) GetImages() {
	var images = []image.Image{}

	for _, v := range *g.imagesAddresses {
		imgFile, err := os.Open(v)

		if err != nil {
			fmt.Println(err)
		}

		img, _, err := image.Decode(imgFile)

		if err != nil {
			fmt.Println(err)
		}

		images = append(images, img)

	}

	g.images = images
}

func (g *GeneratorInputData) GenerateImages() {
	g.GetImages()
	g.ResizeImages()
	g.AddText()
	g.CombineImages()
}

func (g *GeneratorInputData) AddText() {
	_, f, _ := helpers.LoadFont()
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

			helpers.DrawText(c, val, outlineX, outlineY)

			c.SetSrc(image.White)
			c.SetFontSize(fontSize)
			imageX := (rgba.Bounds().Dx() - len(val)*10) / 2
			imageY := j*(rgba.Bounds().Dy()-int(fontSize*1.5)) + int(fontSize)

			helpers.DrawText(c, val, imageX, imageY)

		}

	}
}

func (g *GeneratorInputData) ResizeImages() {
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

	for i, img := range g.images {
		x := img.Bounds().Dx()
		y := img.Bounds().Dy()

		if *g.orientation == "horizontal" {
			var ratio float64 = float64(y) / float64(minHeight)

			x = int(float64(x) / ratio)
			y = minHeight
		} else {
			var ratio float64 = float64(x) / float64(minWidth)

			y = int(float64(y) / ratio)
			x = minWidth
		}

		dst := image.NewRGBA(image.Rect(0, 0, x, y))
		draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

		g.images[i] = dst
	}
}

func (g *GeneratorInputData) CombineImages() {
	var width int = 0
	var height int = 0

	for _, v := range g.images {
		if *g.orientation == "horizontal" {
			width += v.Bounds().Dx()
			height = v.Bounds().Dy()
		} else {
			width = v.Bounds().Dx()
			height += v.Bounds().Dy()
		}
	}

	bigImage := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}
	rgba := image.NewRGBA(bigImage)

	x := 0
	y := 0
	for _, img := range g.images {

		draw.Draw(rgba, image.Rectangle{image.Point{x, y}, image.Point{img.Bounds().Dx() + x, img.Bounds().Dy() + y}}, img, image.Point{0, 0}, draw.Src)

		if *g.orientation == "horizontal" {
			x += img.Bounds().Dx()
		} else {
			y += img.Bounds().Dy()
		}

	}

	helpers.GenerateOutput(rgba)
}
