package pokeapi

func ListLocationAreas(pageURL *string) ([]byte, error) {
	endpoint := "/location-area/"
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