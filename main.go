package main

import (
	"example.com/generator/generator"
)

func main() {
	g := generator.GeneratorInputData{}
	texts := map[int][]string{
		0: []string{"mistakes", "because of dot"},
		1: []string{"", "Developers"},
	}
	imagesSources := []string{"./sources/1.jpeg", "./sources/2.jpeg"}
	orientation := "horizontal"

	g.InitGeneratorValues(&imagesSources, &texts, &orientation)
	g.GenerateImages()
}
