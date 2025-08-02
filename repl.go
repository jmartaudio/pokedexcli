package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
        scanner := bufio.NewScanner(os.Stdin)

        for {
                fmt.Print("Pokedex > ")
                scanner.Scan()

		input := cleanInput(scanner.Text())
                if len(input) == 0 {
                        fmt.Println("Enter some text")
                        continue
                }
		userCmd := input[0]

                cmd, exists := getCommands()[userCmd]
                if exists {
			err := cmd.callback() 
			if err != nil {
				fmt.Println(err)
			}
			continue
                } else {
                        fmt.Printf("Command: %s Not Found\n", userCmd)
			continue
                }
        }
}

func cleanInput(text string) []string {
        inputSlice := strings.Fields(strings.ToLower(text))
        return inputSlice
}

type cliCommand struct {
        name        string
        description string
        callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
