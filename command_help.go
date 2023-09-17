package main

import (
	"fmt"
	"pokemoncli/internal/pokeapi"
)

func commandHelp(config *pokeapi.Client) error {
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s", command.name, command.description)
		fmt.Println()
	}
	return nil
}