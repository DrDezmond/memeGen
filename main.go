package main

import (
	"example.com/generator/generator"
)

func main() {
	g := generator.GeneratorData{}
	texts := map[int][]string{
		0: []string{"mistakes", "because of dot"},
		1: []string{"Fucking", "Developers"},
		2: []string{"Are", "Bastards"},
		3: []string{"Fuck", "This"},
	}
	imagesSources := []string{"./sources/1.jpeg", "./sources/2.jpeg", "./sources/2.jpeg", "./sources/1.jpeg"}
	orientation := "grid"

	g.InitGeneratorValues(&imagesSources, &texts, &orientation)
	g.GenerateImages()
}
