package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/aneesh-mulye/pokedex/internal/pokeapi"
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

		cmdWord := parsedCmd[0]
		cmdFun := commands[cmdWord].callback
		if nil == cmdFun {
			fmt.Fprintln(os.Stderr, "Unknown command")
			continue
		}

		err := cmdFun(parsedCmd[1:])
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

var pokedex map[string]pokeapi.PokemonInfo

func init() {
	pokedex = make(map[string]pokeapi.PokemonInfo)
}

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the next location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore a location and list the pokemon found there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}

func commandExit(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	}

	return nil
}

func commandMap(args []string) error {
	nextLocs, err := pokeapi.NextLocationAreas()
	if err != nil {
		if err.Error() == "END" {
			fmt.Println("you're at the last page")
			return nil
		}
		return err
	}
	for _, loc := range nextLocs {
		fmt.Println(loc)
	}
	return nil
}

func commandMapb(args []string) error {
	previousLocs, err := pokeapi.PreviousLocationAreas()
	if err != nil {
		if err.Error() == "BEGINNING" {
			fmt.Println("you're on the first page")
			return nil
		}
		return err
	}
	for _, loc := range previousLocs {
		fmt.Println(loc)
	}
	return nil
}

func commandExplore(args []string) error {
	if len(args) < 1 {
		fmt.Println("you must name a location to explore")
		return nil
	}

	location := args[0]

	fmt.Println("Exploring " + location + "...")

	pokemonIn, err := pokeapi.PokemonInLocation(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonIn {
		fmt.Println(pokemon)
	}

	return nil
}

func commandCatch(args []string) error {
	if len(args) < 1 {
		fmt.Println("you must name a Pokemon to try to catch")
		return nil
	}

	pokemonName := args[0]

	if _, exists := pokedex[pokemonName]; exists {
		fmt.Println("you've already caught " + pokemonName)
		return nil
	}

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		return err
	}

	caught := rand.Intn(pokemon.BaseExperience) < 50
	if !caught {
		fmt.Println(pokemonName + " escaped!")
		return nil
	}

	pokedex[pokemonName] = pokemon
	fmt.Println(pokemonName + " was caught!")
	fmt.Println("You may now inspect it with the inspect command.")

	return nil
}

func commandInspect(args []string) error {
	if len(args) < 1 {
		fmt.Println("you must name a Pokemon to inspect")
		return nil
	}

	pokemonName := args[0]

	if _, exists := pokedex[pokemonName]; !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	p := pokedex[pokemonName]

	fmt.Println("Name: " + p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  - "+stat.Stat.Name+": %d\n", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Println("  - " + t.Type.Name)
	}

	return nil
}

func commandPokedex(args []string) error {
	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. You have not caught any Pokemon.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemonName := range pokedex {
		fmt.Println(" - " + pokemonName)
	}

	return nil
}
