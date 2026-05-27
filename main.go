package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "Encounter-Manager"

const encounterLogo = `                                                                                                     
  ▄▄▄▄▄▄▄                                                 ▄▄▄     ▄▄▄                                     
 █▀██▀▀▀                                █▄                 ███▄ ▄███                                      
   ██     ▄                       ▄    ▄██▄      ▄         ██ ▀█▀ ██         ▄              ▄▄       ▄    
   ████   ████▄ ▄███▀ ▄███▄ ██ ██ ████▄ ██ ▄█▀█▄ ████▄     ██     ██   ▄▀▀█▄ ████▄ ▄▀▀█▄ ▄████ ▄█▀█▄ ████▄
   ██     ██ ██ ██    ██ ██ ██ ██ ██ ██ ██ ██▄█▀ ██        ██     ██   ▄█▀██ ██ ██ ▄█▀██ ██ ██ ██▄█▀ ██   
   ▀█████▄██ ▀█▄▀███▄▄▀███▀▄▀██▀█▄██ ▀█▄██▄▀█▄▄▄▄█▀      ▀██▀     ▀██▄▄▀█▄██▄██ ▀█▄▀█▄██▄▀████▄▀█▄▄▄▄█▀   
                                                                                            ██            
                                                                                          ▀▀▀             `

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var encounter []Entities
	gameStarted := false
	currentTurn := 0
	fmt.Println(StylePrompt.Render(encounterLogo))
	fmt.Println("Tabletop Encounter Manager Started. Type 'exit' to quit or 'help' to get started")
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
		case "remove":
			encounter, currentTurn = removeEntityCommand(args, encounter, currentTurn)
		case "damage":
			encounter, currentTurn = damageEntity(args, scanner, encounter, currentTurn)
		case "heal":
			encounter = healEntity(args, encounter)
		default:
			invalidCommand()
		}
		printPrompt()
	}
}
