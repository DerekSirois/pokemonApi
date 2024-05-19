package model

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type Pokemon struct {
	Id            int `json:",omitempty"`
	Name          string
	PokedexNumber int
	Type1         string
	Type2         string
}

type PokemonStore struct {
	Db *sqlx.DB
}

func (s *PokemonStore) GetAll() ([]Pokemon, error) {
	pokemon := []Pokemon{}
	err := s.Db.Select(&pokemon, "SELECT p.id, p.name, p.pokedexnumber, t.name as type1, t2.name as type2 FROM pokemon p JOIN types t on p.type1id = t.id JOIN types t2 on p.type2id = t2.id")
	if err != nil {
		return nil, errors.New("failed to get the pokemons")
	}
	return pokemon, nil
}

func (s *PokemonStore) GetById(id int) (Pokemon, error) {
	pokemon := Pokemon{}
	err := s.Db.Get(&pokemon, "SELECT p.id, p.name, p.pokedexnumber, t.name as type1, t2.name as type2 FROM pokemon p JOIN types t on p.type1id = t.id JOIN types t2 on p.type2id = t2.id WHERE p.id = $1", id)
	if err != nil {
		return Pokemon{}, errors.New("failed to get the pokemons")
	}
	return pokemon, nil
}

func (s *PokemonStore) GetRandom() (Pokemon, error) {
	pokemon := Pokemon{}
	err := s.Db.Get(&pokemon, "SELECT p.id, p.name, p.pokedexnumber, t.name as type1, t2.name as type2 FROM pokemon p JOIN types t on p.type1id = t.id JOIN types t2 on p.type2id = t2.id ORDER BY random() LIMIT 1")
	if err != nil {
		return Pokemon{}, errors.New("failed to get the pokemons")
	}
	return pokemon, nil
}

func (s *PokemonStore) GetTypeIdByName(name string) (int, error) {
	pokemonType := struct{ Id int }{}
	err := s.Db.Get(&pokemonType, "SELECT id FROM types WHERE name = $1", name)
	if err != nil {
		return 0, errors.New("failed to create the pokemon")
	}
	return pokemonType.Id, nil
}

func (s *PokemonStore) Create(pokemon Pokemon) error {
	type1, err := s.GetTypeIdByName(pokemon.Type1)
	if err != nil {
		return errors.New("failed to get the type")
	}
	var type2 int
	if pokemon.Type2 != "" {
		type2, err = s.GetTypeIdByName(pokemon.Type2)
		if err != nil {
			return errors.New("failed to get the type")
		}
	}

	// TODO look if their is a better solution
	t2 := (map[bool]any{true: type2, false: nil})[type2 != 0] // hack to have nil if the pokemon doesn't have a second type

	_, err = s.Db.Exec("INSERT INTO pokemon (name, pokedexnumber, type1id, type2id) VALUES ($1, $2, $3, $4)", pokemon.Name, pokemon.PokedexNumber, type1, t2)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
