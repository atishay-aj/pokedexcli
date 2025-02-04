package main

import "github.com/atishay-aj/pokedexcli/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient:           pokeapi.NewClient(),
		nextLocationAreaURL:     nil,
		previousLocationAreaURL: nil,
	}
	startRepl(&cfg)

}
