package main

import "testing"

func FullOneEntityTest(t *testing.T) {
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

func EntityAddedTest(t *testing.T) {
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
			"Object already in",
			[]Entities{{Name: "goblin", Health: 10, Initiative: 5}},
			[]string{"add", "troll", "42", "6"},
			1,
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
