package main

import (
	"go-vk-observer/config"
	"go-vk-observer/internal/app"
	"log"
)

func main() {
	cfg := config.Init()

	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Can't create application: %v", err)
	}

	application.Run()
}
