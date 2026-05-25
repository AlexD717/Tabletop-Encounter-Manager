package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
)

func printPrompt() {
	coloredPrompt := StylePrompt.Render(cliName + "> ")
	fmt.Print(coloredPrompt)
}

func invalidCommand() {
	printError("Invalid command, to list all commands type 'help'")
}

func printCommand(name string, arguments string, description string) {
	leftColumn := StyleCommand.Render(name)

	var rightColumn string
	if len(arguments) != 0 {
		rightColumn = lipgloss.JoinVertical(lipgloss.Left, StyleArguments.Render(arguments), StyleDescription.Render(description))
	} else {
		rightColumn = StyleDescription.Render(description)
	}

	finalRow := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, rightColumn)

	fmt.Printf("%s\n", finalRow)
}

func helpCommand() {
	printCommand("exit", "", "Exits the CLI tool")
	printCommand("list", "", "Lists all entities in the encounter")
	printCommand("current", "", "Lists the entities whose turn it currently is")
	printCommand("next", "", "Ends the current entities turn and says who the next entity is to go")
	printCommand("add", "<name> <health> <initiative>", "Adds a new entity to the encounter")
	printCommand("damage", "<name> <damage-amount>", "Damages the specified entity by the specified amount")
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
