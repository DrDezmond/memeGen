package generatorService

import (
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

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
