package main

import (
	"encoding/json"
	"fmt"
	"pokemoncli/internal/pokeapi"
	"pokemoncli/internal/pokecache"
	"time"
)

var cache pokecache.Cache = pokecache.NewCache(time.Hour)

func commandMap(config *pokeapi.Client) error {
	// declare empty struct for unmarshaling JSON
	locationAreas := pokeapi.LocationAreaResp{}
	// check if this is the first API call
	if config.Next == nil {
		firstReq := pokeapi.BaseURL + "/location"
		config.Next = &firstReq
	} else {
		// check if API call is cached
		dat, ok := cache.Get(*config.Next)
		if ok {
			err := json.Unmarshal(dat, &locationAreas)
			if err != nil {
				return err
			}
			config.Next = locationAreas.Next
			if locationAreas.Previous != nil {
				config.Previous = locationAreas.Previous
			}
			for _, loc := range locationAreas.Results {
				fmt.Printf("%s\n", loc.Name)
			}
			return nil
		}
	}
	// make API call, unmarshal json, add to cache set next
	results, err := pokeapi.ListLocationAreas(config.Next)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	err = json.Unmarshal(results, &locationAreas)
	if err != nil {
		return err
	}
	cache.Add(*config.Next, results)
	config.Next = locationAreas.Next
	if locationAreas.Previous != nil {
		config.Previous = locationAreas.Previous
	}
	for _, loc := range locationAreas.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}

func commandMapB(config *pokeapi.Client) error {
	locationAreas := pokeapi.LocationAreaResp{}
	if config.Previous == nil {
		fmt.Println("No previous locations")
		return nil
	} else {
		dat, ok := cache.Get(*config.Previous)
		if ok {
			err := json.Unmarshal(dat, &locationAreas)
			if err != nil {
				return err
			}
			config.Next = locationAreas.Next
			if locationAreas.Previous != nil {
				config.Previous = locationAreas.Previous
			}
			for _, loc := range locationAreas.Results {
				fmt.Printf("%s\n", loc.Name)
			}
			return nil
		}
	}
	results, err := pokeapi.ListLocationAreas(config.Previous)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	err = json.Unmarshal(results, &locationAreas)
	if err != nil {
		return err
	}

	cache.Add(*config.Previous, results)

	fmt.Println(locationAreas.Previous)
	config.Next = locationAreas.Next
	if (locationAreas.Previous == nil) {
		config.Previous = nil
	} else {
		config.Previous = locationAreas.Previous
	}
	for _, loc := range locationAreas.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}