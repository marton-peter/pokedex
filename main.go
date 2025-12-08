package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Text() == "" {
			continue
		} else {
			text := strings.Fields(strings.TrimSpace(strings.ToLower(scanner.Text())))
			fmt.Printf("Your command was: %s\n", text[0])
		}
	}
}
