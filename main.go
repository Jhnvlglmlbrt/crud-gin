package main

import (
	"log"

	"github.com/Jhnvlglmlbrt/crud-gin/config"
	"github.com/Jhnvlglmlbrt/crud-gin/internal/app"
)

func main() {

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Error at starting: %v", err)
	}

	app.Run(cfg)
}
