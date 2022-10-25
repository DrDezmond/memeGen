package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/DrDezmond/memeGen/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "configs/apiserver.toml", "path to config")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
