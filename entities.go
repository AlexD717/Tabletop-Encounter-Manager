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
		entityString := fmt.Sprintf("%s's", encounter[currentTurn].Name)
		coloredName := StyleEntity.Render(entityString)
		coloredHealth := StyleHealth.Render(fmt.Sprintf("%d", encounter[currentTurn].Health))

		fmt.Printf("It is currently %s turn, and they have %s health\n", coloredName, coloredHealth)
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
		printError("Invalid Number of Arguments. To use add [name] [health] [initiative]")
		return encounter, currentTurn
	}

	name := args[1]
	for _, entity := range encounter {
		if strings.EqualFold(entity.Name, name) {
			printError("Error: There is already an entity with the name %s in the encounter", entity.Name)
			return encounter, currentTurn
		}
	}

	health, err := strconv.Atoi(args[2])
	if err != nil {
		printError("Error: Health must be a number")
		return encounter, currentTurn
	}

	initiative, err := strconv.Atoi(args[3])
	if err != nil {
		printError("Error: Initiative must be a number ")
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

	coloredName := StyleEntity.Render(name)
	coloredHealth := StyleHealth.Render(fmt.Sprintf("%d", health))
	coloredInitiative := StyleInitiative.Render(fmt.Sprintf("%d", initiative))

	fmt.Printf("Added entity %s with %s health and an initiative of %s\n", coloredName, coloredHealth, coloredInitiative)
	return encounter, currentTurn
}

func damageEntity(args []string, scanner *bufio.Scanner, encounter []Entities, currentTurn int) ([]Entities, int) {
	if len(args) != 3 {
		printError("Invalid Number of Arguments. To use damage [name] [amount]")
		return encounter, currentTurn
	}

	damage, err := strconv.Atoi(args[2])
	if err != nil {
		printError("Error: Damage must be a number")
		return encounter, currentTurn
	}

	name := args[1]
	for i, entity := range encounter {
		if strings.EqualFold(entity.Name, name) {
			encounter[i].Health -= damage
			coloredDamage := StyleHealth.Render(fmt.Sprintf("%d", damage))
			coloredHealth := StyleHealth.Render(fmt.Sprintf("%d", encounter[i].Health))
			fmt.Printf("%s damage dealt to %s, new health is %s\n", coloredDamage, StyleEntity.Render(name), coloredHealth)

			if encounter[i].Health <= 0 {
				fmt.Printf("%s health is equal to or below zero, would you like to remove them? (y/n) ", StyleEntity.Render(name))
				if scanner.Scan() {
					answer := cleanInput(scanner.Text())
					if strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes") {
						encounter = append(encounter[:i], encounter[i+1:]...)
						fmt.Printf("removed %s\n", StyleEntity.Render(name))

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

	fmt.Println("No entity found with name: " + StyleEntity.Render(name))
	return encounter, currentTurn
}

func listEntities(encounter []Entities) string {
	if len(encounter) == 0 {
		return "No entities in the current encounter\n"
	}

	var builder strings.Builder
	for _, entity := range encounter {
		coloredHealth := StyleHealth.Render(fmt.Sprintf("%d", entity.Health))
		coloredInitiative := StyleInitiative.Render(fmt.Sprintf("%d", entity.Initiative))
		line := fmt.Sprintf("Entity %s, %s health, %s initiative\n", StyleEntity.Render(entity.Name), coloredHealth, coloredInitiative)
		builder.WriteString(line)
	}

	return builder.String()
}
