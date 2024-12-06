package main

import (
	"time"

	"github.com/bootdotdev/go-api-gate/courses/projects/bootdev_pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Hour)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
