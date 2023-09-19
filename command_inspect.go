package main

import (
	"fmt"
	"pokemoncli/internal/pokeapi"
)

func commandInspect(config *pokeapi.Client, arg string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	// check if pokemon is cached
	dat, ok := pokemonCache.Get(arg)
	if ok {
		fmt.Printf("Name: %s \n", dat.Name)
		fmt.Printf("Height: %d \n", dat.Height)
		fmt.Printf("Weight: %d \n", dat.Weight)
		fmt.Println("Stats:")
		for _, stat := range dat.Stats {
			fmt.Printf("  -%s: %d \n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		//fmt.Printf("%v \n", dat.Types)
		//fmt.Println(len(dat.Types))
		for _, typeStruct := range dat.Types {
			fmt.Printf("  - %s \n", typeStruct.Type.Name)
		}
		return nil
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}