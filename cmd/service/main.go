package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"l0Service/internal/app/service"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/service.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := service.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := service.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
