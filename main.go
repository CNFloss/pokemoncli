package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func trimLastChar(s string) string {
  r, size := utf8.DecodeLastRuneInString(s)
  if r == utf8.RuneError && (size == 0 || size == 1) {
    size = 0
  }
  return s[:len(s)-size]
}

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

func main() {
	for ;; {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Pokedex > ")

    input, _ := reader.ReadString('\n')

		input = trimLastChar(input)
		commands := getCommands()
		command, ok := commands[input]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		command.callback()
	}
}