package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		input_command := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		availableCmds := getCommands()
		command, ok := availableCmds[input_command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback(cfg, args...)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List next location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous location areas",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore {area}",
			description: "List the pokemon in a area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Try to catch a pokemon",
			callback:    callbackCatch,
		},
	}
}
