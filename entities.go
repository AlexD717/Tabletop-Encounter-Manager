package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Entities struct {
	Name       string
	Health     int
	Initiative int
}

func addEntities(args []string) {
	if len(args) != 4 {
		fmt.Println("Invalid Number of Arguments. To use add [name] [health] [initiative]")
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
	fmt.Printf("Added entity %s with %d health and an initiative of %d\n", name, health, initiative)
}

func damageEntity(args []string, scanner *bufio.Scanner) {
	if len(args) != 3 {
		fmt.Println("Invalid Number of Arguments. To use damage [name] [amount]")
		return
	}

	damage, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error: Damage must be a number")
		return
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
					}
				}
			}
			return
		}
	}
	fmt.Println("No entity found with name: " + name)
}

func listEntities() {
	if len(encounter) == 0 {
		fmt.Println("No entities in the current encounter")
	}

	for _, entity := range encounter {
		fmt.Printf("Entity %s, %d health, %d initiative\n", entity.Name, entity.Health, entity.Initiative)
	}
}
