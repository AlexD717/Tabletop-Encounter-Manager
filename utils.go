package main

import (
	"fmt"
	"strings"
)

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
