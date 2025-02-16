package main

import (
	"time"

	"github.com/atishay-aj/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:           pokeapi.NewClient(time.Hour),
		nextLocationAreaURL:     nil,
		previousLocationAreaURL: nil,
		caughtPokemon:           make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)

}
