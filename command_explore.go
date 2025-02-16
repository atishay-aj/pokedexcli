package main

import (
	"fmt"
	"log"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		println("No area provided")
		return nil
	}
	singleAreaName := args[0]
	singleArea, err := cfg.pokeapiClient.GetSingleArea(singleAreaName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pokemon in %s:\n", singleArea.Name)
	for _, pokemon := range singleArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
