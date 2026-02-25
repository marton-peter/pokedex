package main

import (
	"time"

	"github.com/marton-peter/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{client: client}
	startRepl(cfg)
}
