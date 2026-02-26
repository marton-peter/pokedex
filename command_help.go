package main

import (
	"fmt"
)

func commandHelp(cfg *config, args []string) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Displays the names of the next 20 location areas in the Pokemon world
explore: Displays a list of all the Pokémon located in the selected area
`)
	return nil
}
