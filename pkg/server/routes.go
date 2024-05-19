package server

import (
	"pokemonApi/pkg/handler"

	"github.com/go-chi/chi/v5"
)

func (s Server) initRoutes(r *chi.Mux) {
	r.Get("/", handler.Index)
	r.Get("/pokemon", s.pokemonHandler.GetAll)
	r.Get("/pokemon/{id:[0-9]+}", s.pokemonHandler.GetById)
	r.Get("/pokemon/encounter", s.pokemonHandler.Encounter)
	r.Post("/pokemon", s.pokemonHandler.Create)
}
