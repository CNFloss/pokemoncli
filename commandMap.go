package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

func commandMap() error {
	var body []byte
	if len(config.Results) == 0 {
		body = getPokeAPI("https://pokeapi.co/api/v2/location-area/")
	} else if len(config.Next) != 0 {
		body = getPokeAPI(config.Next)
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
		body = getPokeAPI(config.Previous)
	} else {
		return errors.New("No previous locations to get")
	}
	config.jsonToStruct(body)
	for _, loc := range config.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}