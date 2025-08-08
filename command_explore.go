package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name:= args[0]
	locExpResp, err := cfg.pokeapiClient.ExploreLocations(name)
	if err != nil {
		return err
	}
	for _, pokemonEnc := range locExpResp.PokemonEncounters {
		fmt.Println(pokemonEnc.Pokemon.Name)
	}

	return nil
}
