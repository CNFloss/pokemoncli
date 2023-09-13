package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokemoncli/internal/pokecache"
	"time"
)

type jsonMappable interface {
	jsonToStruct([]byte)
}

type Location struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Config struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Location `json:"results"`
}

func (c *Config) jsonToStruct(body []byte) {
	err := json.Unmarshal(body, c)
	if err != nil {
			fmt.Println(err)
	}
}

func getPokeAPI(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	return body
}

var config *Config = &Config{}

var cache pokecache.Cache = pokecache.NewCache(time.Hour)

func commandMap() error {
	var body []byte
	if len(config.Results) == 0 {
		body = getPokeAPI("https://pokeapi.co/api/v2/location-area/")
		cache.Add("https://pokeapi.co/api/v2/location-area/", body)
	} else if len(config.Next) != 0 {
		dat, ok :=  cache.Get(config.Next)
		if ok {
			fmt.Println("cache hit!")
			config.jsonToStruct(dat)
			for _, loc := range config.Results {
				fmt.Printf("%s\n", loc.Name)
			}
			return nil
		}
		fmt.Println("cache miss")
		body = getPokeAPI(config.Next)
		cache.Add(config.Next, body)
	} else {
		return errors.New("No more locations to get")
	}
	config.jsonToStruct(body)
	for _, loc := range config.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}

func commandMapB() error {
	var body []byte
	if len(config.Previous) != 0 {
		fmt.Println("cache hit!")
		dat, ok :=  cache.Get(config.Previous)
		if ok {
			config.jsonToStruct(dat)
			for _, loc := range config.Results {
				fmt.Printf("%s\n", loc.Name)
			}
			return nil
		}
		fmt.Println("cache miss")
		body = getPokeAPI(config.Previous)
		cache.Add(config.Previous, body)
	} else {
		return errors.New("No previous locations to get")
	}
	config.jsonToStruct(body)
	for _, loc := range config.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}