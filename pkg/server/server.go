package server

import (
	"log"
	"net/http"
	"pokemonApi/pkg/database"
	"pokemonApi/pkg/handler"
	"pokemonApi/pkg/model"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	pokemonHandler handler.PokemonHandler
}

func New() *Server {
	db := database.Connect()
	return &Server{
		pokemonHandler: handler.PokemonHandler{Store: &model.PokemonStore{Db: db}},
	}
}

func (s Server) Run() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	s.initRoutes(r)

	log.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}
