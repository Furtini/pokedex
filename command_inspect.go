package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemonName *string) error {
	if pokemonName == nil || *pokemonName == "" {
		return errors.New("You must provide a Pokemon name")
	}

	pokemon, ok := cfg.pokedex[*pokemonName]
	if !ok {
		return errors.New("you have  not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}
