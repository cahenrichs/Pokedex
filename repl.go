package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

type cliCommand struct {
	name		string
	description string
	callback	func() error
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
		
	},
	"map": {
		name:			"map",
		description:	"Displays the locations",
		
	},
	"mapb": {
		name:			"mapb",
		description:	"Goes back a page of locations",
	}
	}
}


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()
	if help, ok := cmds["help"]; ok {
		help.callback = makeHelp(cmds)
		cmds["help"] = help
	}

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		word := cleanInput(input)
		if len(word) == 0 {
			continue
		}

		firstWord := word[0]
		c, ok := cmds[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := c.callback(); err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)

	clean := strings.Fields(lowerText)
	return clean
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cmds map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cmds {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
} 
func makeHelp(cmds map[string]cliCommand) func() error {
	return func() error {return commandHelp(cmds)}
}