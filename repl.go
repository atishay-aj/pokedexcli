package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		command := cleaned[0]
		fmt.Println("Your command was:", command)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
