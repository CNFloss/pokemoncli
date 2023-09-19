package main

import (
	"fmt"
	"pokemoncli/internal/pokeapi"
)

func commandPokedex(config *pokeapi.Client, arg string) error {
	fmt.Println("Your Pokedex:")
	
	for key := range pokemonCache.Cache {
		fmt.Printf(" - %s  \n", key)
	}

	return nil
}