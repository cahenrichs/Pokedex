package main

import "strings"

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)

	clean := strings.Fields(lowerText)
	return clean
}