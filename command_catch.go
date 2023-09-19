package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"pokemoncli/internal/pokeapi"
)

func commandCatch(config *pokeapi.Client, arg string) error {
	pokemon := pokeapi.Pokemon{}
	var chance int
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	// check if API call is cached
	dat, ok := pokemonCache.Get(arg)
	if ok {
		chance = rand.Intn(dat.BaseExperience)
		if chance > dat.BaseExperience / 2 {
			fmt.Printf("%s was caught! \n", arg)
		} else {
			fmt.Printf("%s escaped! \n", arg)
		}
		return nil
	}
	results, err := pokeapi.GetPokemon(arg)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	err = json.Unmarshal(results, &pokemon)
	if err != nil {
		return err
	}
	chance = rand.Intn(pokemon.BaseExperience)
	if chance > pokemon.BaseExperience / 2 {
		fmt.Printf("%s was caught! \n", arg)
		pokemonCache.Add(arg, pokemon)
	} else {
		fmt.Printf("%s escaped! \n", arg)
	}
	return nil
}