package generator

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
)

type GeneratorInputData struct {
	imagesAddresses []string
	texts []string
	images []image.Image
	orientation string
}

func (g GeneratorInputData) Log() {
	fmt.Println("ImagesAdresses: ", g.imagesAddresses)
	fmt.Println("Texts: ", g.texts)
	fmt.Println("Images: ", g.images)
	fmt.Println("Orientation: ", g.orientation)
}

func (g *GeneratorInputData) InitGeneratorValues(addresses, texts []string, orientation string) {
	g.imagesAddresses = addresses
	g.texts = texts
	
	orientations := [4]string{"single", "horizontal", "vertical", "grid"}
	for _, o := range orientations {
		if o == orientation {
			g.orientation = orientation
		}
	}
}

func (g *GeneratorInputData) GetImages() {
	var images = []image.Image{}

	for _, v := range g.imagesAddresses {
		imgFile, err := os.Open(v)		

		if err != nil {
				fmt.Println(err)
		}

		img, _, err := image.Decode(imgFile)
		
		if err != nil {
				fmt.Println(err)
		}

		images = append(images, img);
		
	}

	g.images = images
}

func (g *GeneratorInputData) ResizeHorizontal() {
	var minHeight int = -1

	for _, img := range g.images {
		y := img.Bounds().Dy()
		if (minHeight == -1 || minHeight > y) {
			minHeight = y
		} 
	}

	for i, img := range g.images {
		x := img.Bounds().Dx()
		y := img.Bounds().Dy()

		var ratio float64 = float64(y) / float64(minHeight)

		x = int(float64(x) / ratio)
		y = minHeight

		dst := image.NewRGBA(image.Rect(0, 0, x, y))
		draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)	


		g.images[i] = dst
	}
}

func (g *GeneratorInputData) AddTexToImages() {

}

func (g *GeneratorInputData) GenerateImages() {
	g.GetImages()

	switch g.orientation {
		case "horizontal":
			g.ResizeHorizontal()
			g.CombineImagesHorizontal()	
	}

}

func (g *GeneratorInputData) CombineImagesHorizontal() {
	var width int = 0
	var height int = 0

	for _,v := range g.images {
		width += v.Bounds().Dx()
		height = v.Bounds().Dy()
	}

	bigImage := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}
	rgba := image.NewRGBA(bigImage)

	x := 0
	for _,img := range g.images {
		draw.Draw(rgba, image.Rectangle{image.Point{x, 0}, image.Point{img.Bounds().Dx() + x, img.Bounds().Dy()}}, img, image.Point{0, 0}, draw.Src)
		x += img.Bounds().Dx()
	}

	GenerateOutput(rgba)
}

func GenerateOutput(rgba image.Image) {
	out, err := os.Create("./dist/output.jpg")
	if err != nil {
		fmt.Println(err)
	}
	var opt jpeg.Options
	opt.Quality = 99

	jpeg.Encode(out, rgba, &opt)
}
