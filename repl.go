package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marton-peter/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}

type config struct {
	client           *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Display available commands",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Get the next page of locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Get the previous page of locations",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Get a list of pokemons in the passed location",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Attempts to catch the passed pokemon",
		callback:    commandCatch,
	},
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if scanner.Text() == "" {
			continue
		}

		words := cleanInput(scanner.Text())
		cmdName := words[0]
		args := words[1:]

		cmd, ok := commands[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg, args)
		if err != nil {
			fmt.Println(err)
		}
	}
}
