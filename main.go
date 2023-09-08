package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands()map[string]cliCommand {
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
	}
}

func clearInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
    fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := clearInput(text)
		if (len(cleaned) == 0) {
			continue
		}
		input := cleaned[0]
		
		commands := getCommands()
		command, ok := commands[input]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		command.callback()
	}
}