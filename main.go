package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cliName string = "D&D"
var encounter []Entities

type Entities struct {
	Name       string
	Health     int
	Initiative int
}

func printPrompt() {
	fmt.Print(cliName + "> ")
}

func invalidCommand() {
	fmt.Println("Invalid command, to list all commands type 'help'")
}

func helpCommand() {
	fmt.Println("TODO Implement Help Command")
}

func addEntities(args []string) {
	if len(args) != 4 {
		fmt.Println("Invalid Number of Arguments. Usage add [name] [health] [initiative]")
		return
	}

	name := args[1]
	for _, entity := range encounter {
		if strings.EqualFold(entity.Name, name) {
			fmt.Printf("Error: There is already an entity with the name %s in the encounter\n", entity.Name)
			return
		}
	}

	health, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Health must be a number")
		return
	}

	initiative, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Error: Initiative must be a number ")
		return
	}

	newEntity := Entities{
		Name:       name,
		Health:     health,
		Initiative: initiative,
	}
	encounter = append(encounter, newEntity)
	fmt.Printf("Added entity %s with %d health and an initiative of %v\n", name, health, initiative)
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
		case "add":
			addEntities(args)
		default:
			invalidCommand()
		}
		printPrompt()
	}
}
