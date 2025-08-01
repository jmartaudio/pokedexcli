package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func pntHelp(commands map[string]cliCommand) {
	coms := slices.Collect(maps.Keys(commands))
	fmt.Println("Welcome to the Pokedex!")
	for _, com := range coms {
		fmt.Printf("Usage: %s", commands[com].description)
	}
}

func cleanInput(text string) []string {
	inputSlice := strings.Fields(strings.ToLower(text))
	return inputSlice
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{
	    "exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	    },
	    "help": {
		name:        "help",
		description: "Prints Help Message",
		callback:    pntHelp,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			fmt.Println("Enter some text")
		} else {
			inputSlice := cleanInput(input)
			com, exists := commands[inputSlice[0]]
			if exists {
				com.callback()
			} else {
				fmt.Printf("Command: %s Not Found\n", inputSlice[0])
			}
			
		}
	}
}
