package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	availableCmds := getCommands()
	for _, cmd := range availableCmds {
		fmt.Printf("%v: %v", cmd.name, cmd.description)
		fmt.Println("")
	}

	return nil
}
