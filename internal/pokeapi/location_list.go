package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (ListLocationResult, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		result := ListLocationResult{}
		err := json.Unmarshal(val, &result)
		if err != nil {
			return ListLocationResult{}, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ListLocationResult{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ListLocationResult{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ListLocationResult{}, err
	}

	result := ListLocationResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return ListLocationResult{}, err
	}

	c.cache.Add(url, body)

	return result, nil
}
