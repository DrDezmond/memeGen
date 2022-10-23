package helpers

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
	out, err := os.Create("./dist/output.jpg")
	if err != nil {
		fmt.Println(err)
	}
	var opt jpeg.Options
	opt.Quality = 99

	jpeg.Encode(out, rgba, &opt)
}


func LoadFont() (font.Face, *truetype.Font, error) {
	fontFile := "./fonts/impact.ttf"
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

	return face,f, nil
}

func DrawText(c *freetype.Context, text string, x int, y int) {
	pt := freetype.Pt(x, y)

	c.DrawString(strings.ToUpper(text), pt)
}
