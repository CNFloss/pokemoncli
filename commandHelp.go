package main

import (
	"fmt"
)

func commandHelp() error {
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s", command.name, command.description)
		fmt.Println()
	}
	return nil
}