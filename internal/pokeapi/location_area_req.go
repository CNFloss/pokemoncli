package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func getRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	return body, nil
}

func ListLocationAreas(pageURL *string) ([]byte, error) {
	endpoint := "/location-area"
	fullURL := BaseURL + endpoint
	var res []byte
	if (pageURL != nil) {
		fullURL = *pageURL
	}
	dat, err := getRequest(fullURL)
	if err != nil {
		return res, err
	}

	return dat, nil
}