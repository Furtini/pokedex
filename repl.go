package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Furtini/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	pokedex         map[string]pokeapi.Pokemon
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		userCommand := words[0]
		var location *string
		if len(words) > 1 {
			location = &words[1]
		}

		command, ok := getCommands()[userCommand]
		if !ok {
			fmt.Print("Unknown command\n")
			continue
		}

		err := command.callback(cfg, location)
		if err != nil {
			fmt.Printf("%v", err)
		}

		if err := reader.Err(); err != nil {
			fmt.Printf("Shouldn't see an error scanning a string - %v", err)
		}
	}
}
