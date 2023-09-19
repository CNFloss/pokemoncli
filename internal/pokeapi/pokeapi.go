package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	Next *string
	Previous *string
}

func NewClient() Client {
	return Client {
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *LocationAreaResp) JsonToStruct(body []byte) {
	err := json.Unmarshal(body, c)
	if err != nil {
			fmt.Println(err)
	}
}