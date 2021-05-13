package main

import (
	"github.com/BurntSushi/toml"
	"github.com/pyrta2011/rest-api.git/internal/app/apiserver"
	"log"
)

func main() {
	config := apiserver.NewConfig() // Initialization config file
	_, err := toml.DecodeFile("config/apiserver.toml", config) //decode toml config file
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
