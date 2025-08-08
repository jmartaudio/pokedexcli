package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon")
	}
	caught := make([]string, 0, len(cfg.Pokedex))

	for k := range cfg.Pokedex {
		caught = append(caught, k)
	}
	fmt.Println("Your Pokedex:")
	for _, pokeCaught := range caught {
		fmt.Printf(" -%s\n", pokeCaught)
	}
	
	return nil
}
