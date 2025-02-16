package main

import (
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		println("No pokemon provided")
		return nil
	}
	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Printf("you haven't caught %s\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)

	return nil
}
