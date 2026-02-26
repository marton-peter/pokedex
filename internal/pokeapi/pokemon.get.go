package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemonInfo(name string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	if val, ok := c.cache.Get(url); ok {
		var resp Pokemon
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return Pokemon{}, err
		}
		return resp, err
	}

	res, err := c.client.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, body)
	var resp Pokemon
	err = json.Unmarshal(body, &resp)
	return resp, err
}
