package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) FetchLocationAreas(url string) (LocationAreaResponse, error) {
	if val, ok := c.cache.Get(url); ok {
		var resp LocationAreaResponse
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return resp, err
	}

	res, err := c.client.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return LocationAreaResponse{}, err
	}
	c.cache.Add(url, body)
	var resp LocationAreaResponse
	err = json.Unmarshal(body, &resp)
	return resp, err
}
