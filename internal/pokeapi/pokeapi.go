package pokeapi

import (
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