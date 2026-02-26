package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	// 1. Handle arguments
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemonName := args[0]

	// 2. Fetch the page
	resp, err := cfg.client.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	// 3. Attempt catch
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	threshold := 40
	if rand.Intn(resp.BaseExperience) < threshold {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = resp
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
