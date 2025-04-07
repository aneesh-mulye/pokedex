package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	prompt := "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			break
		}
		rawCmd := scanner.Text()
		parsedCmd := cleanInput(rawCmd)
		if 0 == len(parsedCmd) {
			continue
		}
		firstWord := parsedCmd[0]
		fmt.Println("Your command was: " + firstWord)
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercased := strings.ToLower(trimmed)
	split := strings.Fields(lowercased)

	return split
}
