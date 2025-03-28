package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		result := Pokemon{}
		err := json.Unmarshal(val, &result)
		if err != nil {
			return Pokemon{}, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	result := Pokemon{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	return result, nil
}
