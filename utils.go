package main

import "strings"

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	result := strings.Fields(lowered)
	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List of next location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List of previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all pokemons in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon given its name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show pokemon information if caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all pokemons in the Pokedex",
			callback:    commandPokedex,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
