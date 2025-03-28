package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, location *string) error {
	result, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = &result.Next
	cfg.prevLocationURL = &result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	fmt.Println()

	return nil
}

func commandMapb(cfg *config, location *string) error {
	if cfg.prevLocationURL == nil || *cfg.prevLocationURL == "" {
		return errors.New("You are on the first page\n")
	}

	result, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = &result.Next
	cfg.prevLocationURL = &result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	fmt.Println()

	return nil
}
