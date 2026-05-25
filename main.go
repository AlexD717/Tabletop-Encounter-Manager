package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "Encounter-Manager"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var encounter []Entities
	gameStarted := false
	currentTurn := 0
	fmt.Println("Tabletop Encounter Manager Started. Type 'exit' to quit")
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
			fmt.Print(listEntities(encounter))
		case "current":
			currentEntityTurn(encounter, currentTurn)
		case "next":
			currentTurn = nextTurn(encounter, currentTurn)
			gameStarted = true
		case "add":
			encounter, currentTurn = addEntities(args, encounter, currentTurn, gameStarted)
		case "damage":
			encounter, currentTurn = damageEntity(args, scanner, encounter, currentTurn)
		default:
			invalidCommand()
		}
		printPrompt()
	}
}
