package main

import (
	"customer-counter/internal/app/server"
	"customer-counter/util"
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs", "path to config file")
}

func main() {

	flag.Parse()

	config, err := util.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
