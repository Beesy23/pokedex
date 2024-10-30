package main

import (
	"time"

	"github.com/Beesy23/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
