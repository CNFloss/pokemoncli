package main

import (
	"os"
	"pokemoncli/internal/pokeapi"
)

func commandExit(config *pokeapi.Client, arg string) error {
	os.Exit(0)
	return nil
}
