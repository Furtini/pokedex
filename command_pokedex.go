package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, _ *string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("you have not caught any pokemon\n")
	}

	fmt.Printf("Your Pokedex:\n")
	for key := range cfg.pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
