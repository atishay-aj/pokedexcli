package main

import (
	"fmt"
	"log"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		println("No pokemon provided")
		return nil
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		fmt.Printf("Failed to catch %s!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.caughtPokemon[pokemonName] = pokemon
	return nil
}
