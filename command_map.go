package main

import (
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	// 1. Decide which URL to use
	var url string
	if cfg.nextLocationsURL == nil {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	} else {
		url = *cfg.nextLocationsURL
	}

	// 2. Fetch one page
	resp, err := cfg.client.FetchLocationAreas(url)
	if err != nil {
		return err
	}

	// 3. Print the names
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	// 4. Save next/previous URLs in cfg
	if resp.Next == "" {
		cfg.nextLocationsURL = nil
	} else {
		cfg.nextLocationsURL = &resp.Next
	}

	if resp.Previous == "" {
		cfg.prevLocationsURL = nil
	} else {
		cfg.prevLocationsURL = &resp.Previous
	}

	return nil
}
