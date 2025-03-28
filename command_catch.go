package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, pokemonName *string) error {
	result, err := cfg.pokeapiClient.GetPokemon(*pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", *pokemonName)

	if shouldCatch(result.BaseExperience) {
		fmt.Printf("%v was caught!\n", *pokemonName)
		if _, ok := cfg.pokedex[result.Name]; !ok {
			cfg.pokedex[result.Name] = result
		}
		return nil
	}

	fmt.Printf("%v escaped!\n", *pokemonName)
	return nil
}

func shouldCatch(baseExperience int) bool {
	rand := rand.IntN(baseExperience)

	if rand > 40 {
		return false
	}

	return true
}
