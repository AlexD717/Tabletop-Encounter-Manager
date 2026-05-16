package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "DnD-Encounter"

func printPrompt() {
	fmt.Print(cliName + "> ")
}

func invalidCommand() {
	fmt.Println("Invalid command, to list all commands type 'help'")
}

func helpCommand() {
	fmt.Println("TODO Implement Help Command")
}

func cleanInput(input string) string {
	removedSpace := strings.TrimSpace(input)
	lowercase := strings.ToLower(removedSpace)
	return lowercase
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("D&D Encounter Tracker Started. Type 'exit' to quit")
	printPrompt()

	for reader.Scan() {
		input := cleanInput(reader.Text())
		args := strings.Split(input, " ")
		command := args[0]

		switch command {
		case "help":
			helpCommand()
		case "exit":
			return
		default:
			invalidCommand()
		}
		printPrompt()
	}
}
