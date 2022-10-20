package main

import (
	"example.com/generator/generator"
)

func main() {
	g := generator.GeneratorInputData{}	
	g.InitGeneratorValues([]string{"./sources/1.jpeg", "./sources/2.jpeg"}, []string{"huy", "pizda"}, "horizontal")	
	g.GenerateImages()
	
	
	// g.Log()
}