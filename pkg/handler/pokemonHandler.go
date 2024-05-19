package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokemonApi/pkg/model"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type store interface {
	GetAll() ([]model.Pokemon, error)
	GetById(int) (model.Pokemon, error)
	GetRandom() (model.Pokemon, error)
	Create(model.Pokemon) error
}

type PokemonHandler struct {
	Store store
}

func (h PokemonHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pokemon, err := h.Store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJson(w, http.StatusOK, pokemon)
}

func (h PokemonHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Failed to parse the id", http.StatusBadRequest)
		return
	}

	pokemon, err := h.Store.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJson(w, http.StatusOK, pokemon)
}

func (h PokemonHandler) Create(w http.ResponseWriter, r *http.Request) {
	pokemon := model.Pokemon{}
	err := json.NewDecoder(r.Body).Decode(&pokemon)
	if err != nil {
		http.Error(w, "Failed to parse the pokemon", http.StatusBadRequest)
		return
	}

	err = h.Store.Create(pokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJson(w, http.StatusCreated, nil)
}

func (h PokemonHandler) Encounter(w http.ResponseWriter, r *http.Request) {
	pokemon, err := h.Store.GetRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := struct{ Message string }{
		Message: fmt.Sprintf("You have just encountered a %s", pokemon.Name),
	}

	sendJson(w, http.StatusOK, res)
}
