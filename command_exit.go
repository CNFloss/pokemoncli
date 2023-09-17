package main

import (
	"os"
	"pokemoncli/internal/pokeapi"
)

func commandExit(config *pokeapi.Client) error {
	os.Exit(0)
	return nil
}
