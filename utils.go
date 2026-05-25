package main

import (
	"fmt"
	"os"
	"strings"
)

func printPrompt() {
	coloredPrompt := StylePrompt.Render(cliName + "> ")
	fmt.Print(coloredPrompt)
}

func invalidCommand() {
	printError("Invalid command, to list all commands type 'help'")
}

func helpCommand() {
	fmt.Println("TODO Implement Help Command") // TODO do this
}

func cleanInput(input string) string {
	removedSpace := strings.TrimSpace(input)
	lowercase := strings.ToLower(removedSpace)
	return lowercase
}

func printError(message string, a ...any) {
	errorPrefix := StyleError.Render("Error:")
	message = fmt.Sprintf(message, a...)

	fmt.Fprintf(os.Stderr, "%s %s\n", errorPrefix, message)
}
