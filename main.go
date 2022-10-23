package main

import (
	"example.com/generator/generator"
)

func main() {
	g := generator.GeneratorInputData{}	
	m := make(map[int][]string)
	m[0] = []string{"mistake", "because of dot"}
	m[1] = []string{"", "Developers"}
	g.InitGeneratorValues([]string{"./sources/1.jpeg", "./sources/2.jpeg"}, m, "horizontal")	
	g.GenerateImages()
}

