package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	// 1. Handle arguments
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemonName := args[0]

	// 2. Fetch the info from the pokedex
	subject, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", subject.Name)
	fmt.Printf("Height: %v\n", subject.Height)
	fmt.Printf("Weight: %v\n", subject.Weight)

	fmt.Println("Stats:")
	for _, stat := range subject.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeInfo := range subject.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}
