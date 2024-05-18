package main

import "pokemonApi/pkg/server"

func main() {
	s := server.New()
	s.Run()
}
