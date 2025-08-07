package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	locExpResp, err := cfg.pokeapiClient.ExploreLocations(cfg.argument)
	if err != nil {
		return err
	}
	for _, pokemonEnc := range locExpResp.PokemonEncounters {
		fmt.Println(pokemonEnc.Pokemon.Name)
	}

	return nil
}
