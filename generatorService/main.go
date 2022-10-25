package main

import (
	generatorService "github.com/DrDezmond/memeGen/generatorService/generator"
)

func main() {
	g := generatorService.GeneratorData{}
	texts := map[int][]string{
		0: []string{"mistakes", "because of dot"},
		1: []string{"Fucking", "Developers"},
		2: []string{"Are", "Bastards"},
		3: []string{"Fuck", "This"},
	}
	imagesSources := []string{"./generatorService/sources/1.jpeg", "./generatorService/sources/2.jpeg", "./generatorService/sources/2.jpeg", "./generatorService/sources/1.jpeg"}
	orientation := "horizontal"

	g.InitGeneratorValues(&imagesSources, &texts, &orientation)
	g.GenerateImages()
}
