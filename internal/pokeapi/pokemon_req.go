package pokeapi

func GetPokemon(pokemon string) ([]byte, error) {
	var res []byte
	endpoint := "/pokemon/"
	fullURL := BaseURL + endpoint + pokemon
	dat, err := getRequest(fullURL)
	if err != nil {
		return res, err
	}

	return dat, nil
}