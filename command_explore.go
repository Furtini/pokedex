package main

import (
	"fmt"
)

func commandExplore(cfg *config, location *string) error {
	result, err := cfg.pokeapiClient.GetLocationArea(*location)
	if err != nil {
		return err
	}

	for _, encounter := range result.PokemonEntountes {
		fmt.Println(encounter.Pokemon.Name)
	}

	fmt.Println()

	return nil
}
