package main

import (
	"log"
	"pokemonApi/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := server.New()
	s.Run()
}
