package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(locationName string) (LocationAreaResult, error) {
	url := baseUrl + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		result := LocationAreaResult{}
		err := json.Unmarshal(val, &result)
		if err != nil {
			return LocationAreaResult{}, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResult{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResult{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResult{}, err
	}

	result := LocationAreaResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return LocationAreaResult{}, err
	}

	c.cache.Add(url, body)

	return result, nil
}
