package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type LocationAreaResp struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []Location `json:"results"`
}

func (c *LocationAreaResp) JsonToStruct(body []byte) {
	err := json.Unmarshal(body, c)
	if err != nil {
			fmt.Println(err)
	}
}