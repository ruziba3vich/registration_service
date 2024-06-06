package main

import (
	"log"

	"github.com/ruziba3vich/registration_ms/api"
	"github.com/ruziba3vich/registration_ms/internal/config"
	"github.com/ruziba3vich/registration_ms/internal/storage"
)

func main() {
	var config config.Config
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	api := api.New(storage.NewStorage(&config))

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
