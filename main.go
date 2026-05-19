package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "D&D"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var encounter []Entities
	fmt.Println("D&D Encounter Tracker Started. Type 'exit' to quit")
	printPrompt()

	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		args := strings.Split(input, " ")
		command := args[0]

		switch command {
		case "help":
			helpCommand()
		case "exit":
			return
		case "list":
			listEntities(encounter)
		case "add":
			encounter = addEntities(args, encounter)
		case "damage":
			encounter = damageEntity(args, scanner, encounter)
		default:
			invalidCommand()
		}
		printPrompt()
	}
}
