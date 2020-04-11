package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/isratmir/restapi/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf-path", "configs/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.NewAPIServer(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
