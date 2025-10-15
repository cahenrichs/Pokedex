package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		word := cleanInput(input)
		firstWord := word[0]
		if len(word) > 0 {
			fmt.Printf("Your command was: %s\n", firstWord)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)

	clean := strings.Fields(lowerText)
	return clean
}