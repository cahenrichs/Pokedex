package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"github.com/cahenrichs/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name		string
	description string
	callback	func(*config, ...string) error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
	"exit": {
		name:			"exit",
		description:	"Exit the Pokedex",
		callback:		commandExit,
	},
	"help": {
		name:			"help",
		description:	"Displays a help message",
		callback:		commandHelp,
		
	},
	"map": {
		name:			"map",
		description:	"Get the next page of locations",
		callback:		commandMap,
		
	},
	"mapb": {
		name:			"mapb",
		description:	"Get the previous page of locations",
		callback:		commandMapb,
	},
	"explore": {
		name:			"explore <location name>",
		description:	"Explore a location",
		callback:		commandExplore,
	},
	}
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

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)

	clean := strings.Fields(lowerText)
	return clean
}




//func makeHelp(cmds map[string]cliCommand) func() error {
	//return func() error {return commandHelp(cmds)}
//}