package main

import (
	"bufio"
	"fmt"
	"os"
)

var cmdList map[string]cliCommand

func init() {
	cmdList = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Shows all commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Gokedex",
			callback:    commandExit,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func commandHelp() {
	for _, cmd := range cmdList {
		fmt.Printf("%s:\t%s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
}

func commandExit() {
	os.Exit(0)
}

func main() {
	fmt.Println("Welcome to the Gokedex!\nUsage:")
	commandHelp()
	
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Gokedex> ")
		input, err := reader.ReadString('\n')
		input = input[:len(input)-2]

		if err != nil {
			fmt.Printf("Error occured while reading input\n\n")
			continue
		}
		
		if cmd, ok := cmdList[input]; ok {
			cmd.callback()
		} else {
			fmt.Printf("Error: %s not a valid command\n\n", input)
		}
	}

}
