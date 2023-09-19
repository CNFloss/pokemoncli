package main

import (
	"encoding/json"
	"fmt"
	"pokemoncli/internal/pokeapi"
)

func commandExplore(config *pokeapi.Client, arg string) error {
	pokemon := pokeapi.LocationDetailsResp{}
	// check if API call is cached
	dat, ok := cache.Get(arg)
	if ok {
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return err
		}
		for _, loc := range pokemon.PokemonEncounters {
			fmt.Printf("%s\n", loc.Pokemon.Name)
		}
		return nil
	}
	results, err := pokeapi.ListPokemonEncounters(arg)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	cache.Add(arg, results)
	err = json.Unmarshal(results, &pokemon)
	if err != nil {
		return err
	}
	for _, loc := range pokemon.PokemonEncounters {
		fmt.Printf("%s\n", loc.Pokemon.Name)
	}
	return nil
}