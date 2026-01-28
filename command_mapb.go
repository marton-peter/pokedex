package main

import (
	"fmt"

	"github.com/marton-peter/pokedex/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	// 1. Decide which URL to use
	var url string
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = *cfg.prevLocationsURL
	}

	// 2. Fetch one page
	resp, err := pokeapi.FetchLocationAreas(url)
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
