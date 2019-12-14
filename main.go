package main

import (
	"log"
	"test/internal/app/api"
)

func main() {
	log.Println("Starting server")
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
