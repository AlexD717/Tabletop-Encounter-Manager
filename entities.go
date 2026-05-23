package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entities struct {
	Name       string
	Health     int
	Initiative int
}

func sortEntities(encounter []Entities) []Entities {
	sort.SliceStable(encounter, func(i, j int) bool {
		return encounter[i].Initiative > encounter[j].Initiative
	})

	return encounter
}

func currentEntityTurn(encounter []Entities, currentTurn int) {
	if len(encounter) > 0 {
		fmt.Printf("It is currently %s's turn, and they have %d health\n", encounter[currentTurn].Name, encounter[currentTurn].Health)
	} else {
		fmt.Println("There are no entities in the encounter")
	}
}

func nextTurn(encounter []Entities, currentTurn int) int {
	currentTurn = (currentTurn + 1) % len(encounter)
	currentEntityTurn(encounter, currentTurn)
	return currentTurn
}

func addEntities(args []string, encounter []Entities, currentTurn int, gameStarted bool) ([]Entities, int) {
	if len(args) != 4 {
		fmt.Println("Invalid Number of Arguments. To use add [name] [health] [initiative]")
		return encounter, currentTurn
	}

	name := args[1]
	for _, entity := range encounter {
		if strings.EqualFold(entity.Name, name) {
			fmt.Printf("Error: There is already an entity with the name %s in the encounter\n", entity.Name)
			return encounter, currentTurn
		}
	}

	health, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Health must be a number")
		return encounter, currentTurn
	}

	initiative, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Error: Initiative must be a number ")
		return encounter, currentTurn
	}

	newEntity := Entities{
		Name:       name,
		Health:     health,
		Initiative: initiative,
	}

	var activeCharacterName string
	if len(encounter) > 0 {
		activeCharacterName = encounter[currentTurn].Name
	}

	encounter = sortEntities(append(encounter, newEntity))

	if gameStarted && activeCharacterName != "" {
		for i, entity := range encounter {
			if entity.Name == activeCharacterName {
				currentTurn = i
				break
			}
		}
	}

	fmt.Printf("Added entity %s with %d health and an initiative of %d\n", name, health, initiative)
	return encounter, currentTurn
}

func damageEntity(args []string, scanner *bufio.Scanner, encounter []Entities, currentTurn int) ([]Entities, int) {
	if len(args) != 3 {
		fmt.Println("Invalid Number of Arguments. To use damage [name] [amount]")
		return encounter, currentTurn
	}

	damage, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Damage must be a number")
		return encounter, currentTurn
	}

	name := args[1]
	for i, entity := range encounter {
		if strings.EqualFold(entity.Name, name) {
			encounter[i].Health -= damage
			fmt.Printf("%d damage dealt to %s, new health is %d\n", damage, name, encounter[i].Health)

			if encounter[i].Health <= 0 {
				fmt.Printf("%s health is equal to or below zero, would you like to remove them? (y/n) ", name)
				if scanner.Scan() {
					answer := cleanInput(scanner.Text())
					fmt.Printf("\n")

					if strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes") {
						encounter = append(encounter[:i], encounter[i+1:]...)
						fmt.Printf("removed %s\n", name)

						if i < currentTurn {
							currentTurn--
						} else if i == currentTurn {
							if currentTurn >= len(encounter) {
								currentTurn = 0
							}
						}
					}
				}
			}

			return encounter, currentTurn
		}
	}

	fmt.Println("No entity found with name: " + name)
	return encounter, currentTurn
}

func listEntities(encounter []Entities) string {
	if len(encounter) == 0 {
		return "No entities in the current encounter\n"
	}

	var builder strings.Builder
	for _, entity := range encounter {
		line := fmt.Sprintf("Entity %s, %d health, %d initiative\n", entity.Name, entity.Health, entity.Initiative)
		builder.WriteString(line)
	}

	return builder.String()
}
