package main

import (
	"bufio"
	"fmt"
	"os"
	"pokemoncli/internal/pokeapi"
	"pokemoncli/internal/pokecache"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Client, string) error
}

func getCommands()map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Gets 20 locations areas in the Pokemon world, subsequent calls get the next 20 until the end of list",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous 20 locations areas in the Pokemon world, if not at beginning list",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "takes an area as an argument and returns it's details",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch the pokemon provided as an argument",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "shows detail about a caught pokemon provided as an argument",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "shows all the caught pokemon",
			callback:    commandPokedex,
		},
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

var cache pokecache.Cache = pokecache.NewCache(time.Hour)
var pokemonCache pokecache.PokemonCache = pokecache.NewPokemonCache(time.Hour)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &pokeapi.Client{}
	for ;; {
    fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := clearInput(text)
		if (len(cleaned) == 0) {
			continue
		}
		input := cleaned[0]
		arg := ""
		if len(cleaned) > 1 {
			arg = cleaned[1]
		}
		
		commands := getCommands()
		command, ok := commands[input]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		if len(arg) > 0 && !(command.name == "explore" || command.name == "catch" || command.name == "inspect") {
			fmt.Println("this command does not take arguments")
			continue
		}
		command.callback(config, arg)
	}
}