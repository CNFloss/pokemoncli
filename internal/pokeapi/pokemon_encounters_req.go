package pokeapi

func ListPokemonEncounters(area string) ([]byte, error) {
	var res []byte
	endpoint := "/location-area/"
	fullURL := BaseURL + endpoint + area
	dat, err := getRequest(fullURL)
	if err != nil {
		return res, err
	}

	return dat, nil
}