package main

import (
	"fmt"
	"pokemoncli/internal/pokeapi"
)

func commandHelp(config *pokeapi.Client, arg string) error {
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s", command.name, command.description)
		fmt.Println()
	}
	return nil
}