package main

import (
	"fmt"
	"errors"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name:= args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n\n", name)
	pokemonResp, err := cfg.pokeapiClient.FetchPokemonInfo(name)
	if err != nil {
		return err
	}
	randomNumber := rand.IntN(pokemonResp.BaseExperience)
	if randomNumber > 40 {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}

	fmt.Printf("You Caught %s\n", name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.Pokedex[name] = pokemonResp
	return nil
}
