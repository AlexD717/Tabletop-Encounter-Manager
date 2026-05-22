package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestOneFullEntity(t *testing.T) {
	t.Parallel()

	emptyEncounter := []Entities{}
	args := []string{"add", "goblin", "10", "5"}

	resultingEncounter := addEntities(args, emptyEncounter)

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
	}{
		{
			"Valid Addition",
			[]Entities{},
			[]string{"add", "goblin", "10", "7"},
			1,
		},
		{
			"Not Enough Arguments",
			[]Entities{},
			[]string{"add", "goblin", "10"},
			0,
		},
		{
			"Too Many Arguments",
			[]Entities{},
			[]string{"add", "goblin", "10", "7", "15"},
			0,
		},
		{
			"Wrong Health Type",
			[]Entities{},
			[]string{"add", "goblin", "health", "7"},
			0,
		},
		{
			"Wrong Initiative Type",
			[]Entities{},
			[]string{"add", "goblin", "4", "initiative"},
			0,
		},
		{
			"Multiple Wrong Types",
			[]Entities{},
			[]string{"add", "goblin", "health", "initiative"},
			0,
		},
		{
			"Duplicate Name",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 5}},
			[]string{"add", "goblin", "4", "6"},
			1,
		},
		{
			"One Object Already Inside",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 5}},
			[]string{"add", "troll", "42", "6"},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := addEntities(tt.args, tt.initialEncounter)

			if len(result) != tt.expectedLength {
				t.Errorf("Expected encounter length %d, got %d", tt.expectedLength, len(result))
			}
		})
	}
}

func TestDamageEntity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		encounter          []Entities
		args               []string
		fakeUserInput      string
		resultingEncounter []Entities
	}{
		{
			"Valid Damage",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "8"},
			"",
			[]Entities{{Name: "goblin", Health: 2, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
		},
		{
			"Remove Enemy (y)",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "10"},
			"y",
			[]Entities{{Name: "troll", Health: 20, Initiative: 10}},
		},
		{
			"Remove Enemy (yes)",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "15"},
			"y",
			[]Entities{{Name: "troll", Health: 20, Initiative: 10}},
		},
		{
			"Not Remove Enemy (n)",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "15"},
			"n",
			[]Entities{{Name: "goblin", Health: -5, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
		},
		{
			"Not Remove Enemy (invalid input)",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "10"},
			"Random letters",
			[]Entities{{Name: "goblin", Health: 0, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
		},
		{
			"Empty Encounter",
			[]Entities{},
			[]string{"damage", "goblin", "20"},
			"",
			[]Entities{},
		},
		{
			"Too Little Arguments",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin"},
			"",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
		},
		{
			"Too Many Arguments",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "8", "10"},
			"",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
		},
		{
			"Invalid Damage Type",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
			[]string{"damage", "goblin", "number"},
			"",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 4}, {Name: "troll", Health: 20, Initiative: 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fakeInput := bufio.NewScanner(strings.NewReader(tt.fakeUserInput))
			result := damageEntity(tt.args, fakeInput, tt.encounter)

			if !reflect.DeepEqual(result, tt.resultingEncounter) {
				t.Errorf("Expected encounter to be %+v, got %+v", tt.resultingEncounter, result)
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
