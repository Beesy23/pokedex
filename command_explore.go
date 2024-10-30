package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocations(&locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s", locationArea.Name)
	for _, pokemon_encounter := range locationArea.PokemonEncounters {
		fmt.Println(pokemon_encounter.Pokemon.Name)
	}
	return nil
}
