package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestNextEntity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		encounter    []Entities
		currentTurn  int
		expectedTurn int
	}{
		{
			"No Entities in Encounter",
			[]Entities{},
			0,
			0,
		},
		{
			"Next Turn",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}, {Name: "dragon", Health: 50, Initiative: 5}},
			1,
			2,
		},
		{
			"Loop Around",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}, {Name: "dragon", Health: 50, Initiative: 5}},
			2,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resultingTurn := nextTurn(tt.encounter, tt.currentTurn)
			if resultingTurn != tt.expectedTurn {
				t.Errorf("Expected current turn to be %d, got %d", tt.expectedTurn, resultingTurn)
			}
		})
	}
}

func TestOneFullEntity(t *testing.T) {
	t.Parallel()

	emptyEncounter := []Entities{}
	args := []string{"add", "goblin", "10", "5"}

	resultingEncounter, _ := addEntities(args, emptyEncounter, 0, false)

	if len(resultingEncounter) != 1 {
		t.Fatalf("Expected one entity, had %d", len(resultingEncounter))
	}

	if resultingEncounter[0].Name != "goblin" {
		t.Errorf("Expected name 'goblin', got %s", resultingEncounter[0].Name)
	}

	if resultingEncounter[0].Health != 10 {
		t.Errorf("Expected 10 health, got %d", resultingEncounter[0].Health)
	}

	if resultingEncounter[0].Initiative != 5 {
		t.Errorf("Expected initiative of 5, got %d", resultingEncounter[0].Initiative)
	}
}

func TestAddEntity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		initialEncounter []Entities
		args             []string
		expectedLength   int
		currentTurn      int
		expectedTurn     int
		gameStarted      bool
	}{
		{
			"Valid Addition",
			[]Entities{},
			[]string{"add", "goblin", "10", "7"},
			1,
			0,
			0,
			false,
		},
		{
			"Not Enough Arguments",
			[]Entities{},
			[]string{"add", "goblin", "10"},
			0,
			0,
			0,
			false,
		},
		{
			"Too Many Arguments",
			[]Entities{},
			[]string{"add", "goblin", "10", "7", "15"},
			0,
			0,
			0,
			false,
		},
		{
			"Wrong Health Type",
			[]Entities{},
			[]string{"add", "goblin", "health", "7"},
			0,
			0,
			0,
			false,
		},
		{
			"Wrong Initiative Type",
			[]Entities{},
			[]string{"add", "goblin", "4", "initiative"},
			0,
			0,
			0,
			false,
		},
		{
			"Multiple Wrong Types",
			[]Entities{},
			[]string{"add", "goblin", "health", "initiative"},
			0,
			0,
			0,
			false,
		},
		{
			"Duplicate Name",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 5}},
			[]string{"add", "goblin", "4", "6"},
			1,
			0,
			0,
			false,
		},
		{
			"One Object Already Inside Game Not Started",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 5}},
			[]string{"add", "troll", "42", "6"},
			2,
			0,
			0,
			false,
		},
		{
			"Game Started, Higher Initiative",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 5}},
			[]string{"add", "troll", "42", "6"},
			2,
			0,
			1,
			true,
		},
		{
			"Game Started, Lower Initiative",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 5}},
			[]string{"add", "dragon", "50", "1"},
			2,
			0,
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resultingEncounter, resultingTurn := addEntities(tt.args, tt.initialEncounter, tt.currentTurn, tt.gameStarted)

			if len(resultingEncounter) != tt.expectedLength {
				t.Errorf("Expected encounter length %d, got %d", tt.expectedLength, len(resultingEncounter))
			}
			if resultingTurn != tt.expectedTurn {
				t.Errorf("Expected current turn to be %d, got %d", tt.expectedTurn, resultingTurn)
			}
		})
	}
}

func TestDamageEntity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		encounter         []Entities
		args              []string
		fakeUserInput     string
		expectedEncounter []Entities
		currentTurn       int
		expectedTurn      int
	}{
		{
			"Valid Damage",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "8"},
			"",
			[]Entities{{Name: "goblin", Health: 2, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}},
			0,
			0,
		},
		{
			"Last Character Dies",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "troll", "50"},
			"y",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}},
			1,
			0,
		},
		{
			"Current Character Dies",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "15"},
			"y",
			[]Entities{{Name: "troll", Health: 20, Initiative: 10}},
			0,
			0,
		},
		{
			"Character Before Active Dies",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 14}, {Name: "troll", Health: 20, Initiative: 10}, {Name: "dragon", Health: 50, Initiative: 5}},
			[]string{"damage", "goblin", "15"},
			"y",
			[]Entities{{Name: "troll", Health: 20, Initiative: 10}, {Name: "dragon", Health: 50, Initiative: 5}},
			2,
			1,
		},
		{
			"Not Remove Enemy (n)",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "15"},
			"n",
			[]Entities{{Name: "goblin", Health: -5, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			0,
			0,
		},
		{
			"Not Remove Enemy (invalid input)",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "10"},
			"Random letters",
			[]Entities{{Name: "goblin", Health: 0, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			0,
			0,
		},
		{
			"Empty Encounter",
			[]Entities{},
			[]string{"damage", "goblin", "20"},
			"",
			[]Entities{},
			0,
			0,
		},
		{
			"Too Little Arguments",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin"},
			"",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			0,
			0,
		},
		{
			"Too Many Arguments",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "8", "10"},
			"",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			0,
			0,
		},
		{
			"Invalid Damage Type",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "number"},
			"",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			0,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fakeInput := bufio.NewScanner(strings.NewReader(tt.fakeUserInput))
			resultingEncounter, resultingTurn := damageEntity(tt.args, fakeInput, tt.encounter, tt.currentTurn)

			if !reflect.DeepEqual(resultingEncounter, tt.expectedEncounter) {
				t.Errorf("Expected encounter to be %+v, got %+v", tt.expectedEncounter, resultingEncounter)
			}
			if resultingTurn != tt.expectedTurn {
				t.Errorf("Expected current turn to be %d, got %d", tt.currentTurn, tt.expectedTurn)
			}
		})
	}
}

func TestListEntities(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		encounter       []Entities
		expectedStrings []string
	}{
		{
			"Empty Encounter",
			[]Entities{},
			[]string{"no entities"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := listEntities(tt.encounter)
			lowerResult := strings.ToLower(result)

			for _, expected := range tt.expectedStrings {
				if !strings.Contains(lowerResult, expected) {
					t.Errorf("Expected result to contain '%s', but it did not. Result was: %s", expected, result)
				}
			}
		})
	}
}
