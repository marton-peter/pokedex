package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	if val, ok := c.cache.Get(url); ok {
		var resp LocationArea
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return LocationArea{}, err
		}
		return resp, err
	}

	res, err := c.client.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, body)
	var resp LocationArea
	err = json.Unmarshal(body, &resp)
	return resp, err
}
