package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationEncounters, err := cfg.pokeapiClient.GetLocation(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:", locationEncounters.Name)
	for _, pokemon := range locationEncounters.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
