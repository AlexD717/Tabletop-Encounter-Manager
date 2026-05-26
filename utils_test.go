package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Normal Input", "add goblin", "add goblin"},
		{"Uppercase", "ADD GOBLIN", "add goblin"},
		{"Trailing spaces", "   add goblin  	", "add goblin"},
		{"Mixed case and spaces", "ADD GobLiN", "add goblin"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := cleanInput(tt.input)
			if result != tt.expected {
				t.Errorf("cleanInput failed. Expected '%s' but got '%s'", tt.expected, result)
			}
		})
	}
}

func TestRemoveEntityByIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		encounter         []Entities
		index             int
		currentTurn       int
		expectedEncounter []Entities
		expectedTurn      int
	}{
		{
			"Removing Entity Before Active Entity",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}, {Name: "dragon", Health: 200, Initiative: 10}},
			0,
			1,
			[]Entities{{Name: "troll", Health: 20, Initiative: 10}, {Name: "dragon", Health: 200, Initiative: 10}},
			0,
		},
		{
			"Removing Currently Active Entity",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}, {Name: "dragon", Health: 200, Initiative: 10}},
			2,
			2,
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}},
			0,
		},
		{
			"Removing Entity After Active Entity",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}, {Name: "dragon", Health: 200, Initiative: 10}},
			1,
			0,
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "dragon", Health: 200, Initiative: 10}},
			0,
		},
		{
			"Index out of Range",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}},
			420,
			1,
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resultingEncounter, resultingTurn := removeEntityByIndex(tt.index, tt.encounter, tt.currentTurn)

			if !reflect.DeepEqual(resultingEncounter, tt.expectedEncounter) {
				t.Errorf("Expected encounter to be %+v, got %+v", tt.expectedEncounter, resultingEncounter)
			}

			if resultingTurn != tt.expectedTurn {
				t.Errorf("Expected current turn to be %d, got %d", tt.currentTurn, tt.expectedTurn)
			}
		})
	}
}
