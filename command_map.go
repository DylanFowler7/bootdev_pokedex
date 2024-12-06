package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.GetLocationArea(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("no previous locations")
	}
	locationResp, err := cfg.pokeapiClient.GetLocationArea(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
