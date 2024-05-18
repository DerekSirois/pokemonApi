package server

import (
	"pokemonApi/pkg/handler"

	"github.com/go-chi/chi/v5"
)

func initRoutes(r *chi.Mux) {
	r.Get("/", handler.Index)
}
