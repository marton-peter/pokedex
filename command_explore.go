package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	// 1. Handle arguments
	if len(args) == 0 {
		return fmt.Errorf("please provide a location name")
	}
	locationName := args[0]

	// 2. Fetch the page
	resp, err := cfg.client.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	// 3. Print the names
	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
