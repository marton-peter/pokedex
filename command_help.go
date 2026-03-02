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
mapb: Displays the names of the previous 20 location areas in the Pokemon world
explore: Displays a list of all the Pokémon located in the selected area
catch: Attempts to catch the selected pokemon
inspect: Prints information about the selected caught pokemon
pokedex: Lists all caught pokemon
`)
	return nil
}
