package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	initialiseCommandRegistry()

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

		cmdWord := parsedCmd[0]
		cmdFun := commands[cmdWord].callback
		if nil == cmdFun {
			fmt.Fprintln(os.Stderr, "Unknown command")
			continue
		}

		err := cmdFun()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			continue
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercased := strings.ToLower(trimmed)
	split := strings.Fields(lowercased)

	return split
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    nil,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    nil,
	},
}

var commandCallbacks = map[string]func() error{
	"exit": commandExit,
	"help": commandHelp,
}

func initialiseCommandRegistry() {
	for n, cmd := range commands {
		wCallback := cliCommand{
			name:        cmd.name,
			description: cmd.description,
			callback:    commandCallbacks[n],
		}
		commands[n] = wCallback
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	}

	return nil
}
