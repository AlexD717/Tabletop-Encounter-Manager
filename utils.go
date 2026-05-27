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
	printCommand("next", "", "Ends the current entities turn and displays the next entity")
	printCommand("add", "<name> <health> <initiative>", "Adds a new entity to the encounter")
	printCommand("remove", "<name>", "Removes the specified entity from the encounter")
	printCommand("damage", "<name> <damage-amount>", "Damages the specified entity by the specified amount")
	printCommand("heal", "<name> <heal-amount>", "Heals the specified entity by the specified amount")
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

func removeEntityByIndex(index int, encounter []Entities, currentTurn int) ([]Entities, int) {
	if index >= len(encounter) {
		printError("Trying to remove entity at position %d, but the encounter has only %d entities", index, len(encounter))
		return encounter, currentTurn
	}

	entityName := encounter[index].Name
	encounter = append(encounter[:index], encounter[index+1:]...)
	fmt.Printf("Removed %s from the encounter\n", StyleEntity.Render(entityName))

	if index < currentTurn {
		currentTurn--
	} else if index == currentTurn {
		if currentTurn >= len(encounter) {
			currentTurn = 0
		}
	}

	return encounter, currentTurn
}
