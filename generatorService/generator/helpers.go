package generatorService

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func GenerateOutput(rgba image.Image) {
	out, err := os.Create("./generatorService/dist/output.jpg")
	if err != nil {
		fmt.Println(err)
	}
	var opt jpeg.Options
	opt.Quality = 99

	jpeg.Encode(out, rgba, &opt)
}

func LoadFont() (font.Face, *truetype.Font, error) {
	fontFile := "./generatorService/fonts/impact.ttf"
	fontBytes, err := os.ReadFile(fontFile)

	if err != nil {
		return nil, nil, err
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, nil, err
	}
	opts := truetype.Options{}
	opts.Size = 24.0
	face := truetype.NewFace(f, &opts)

	return face, f, nil
}

func DrawText(c *freetype.Context, text string, x int, y int) {
	pt := freetype.Pt(x, y)

	c.DrawString(strings.ToUpper(text), pt)
}

func OpenImages(imagesAddresses []string) []image.Image {
	var images = []image.Image{}

	for _, v := range imagesAddresses {
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
	return images
}
