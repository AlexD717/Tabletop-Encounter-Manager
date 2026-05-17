package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "D&D"
var encounter []Entities

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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
			listEntities()
		case "add":
			addEntities(args)
		case "damage":
			damageEntity(args, scanner)
		default:
			invalidCommand()
		}
		printPrompt()
	}
}
