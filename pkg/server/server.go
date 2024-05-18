package server

import (
	"log"
	"net/http"
	"pokemonApi/pkg/database"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	Db *sqlx.DB
}

func New() *Server {
	return &Server{
		Db: database.Connect(),
	}
}

func (s Server) Run() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	initRoutes(r)

	log.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}
