package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if scanner.Text() == "" {
			continue
		}

		words := cleanInput(scanner.Text())
		cmdName := words[0]

		cmd, ok := commands[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
