package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestFullEncounter(t *testing.T) {
	var encounter []Entities
	currentTurn := 0

	encounter, currentTurn = addEntities([]string{"add", "goblin", "10", "8"}, encounter, currentTurn, false)
	encounter, currentTurn = addEntities([]string{"add", "alex", "10", "20"}, encounter, currentTurn, false)

	if encounter[0].Name != "alex" || encounter[1].Name != "goblin" {
		t.Fatalf("Setup failed, order is wrong")
	}

	fakeInput := bufio.NewScanner(strings.NewReader("n\n"))
	encounter, currentTurn = damageEntity([]string{"damage", "goblin", "7"}, fakeInput, encounter, currentTurn)

	if encounter[1].Health != 3 {
		t.Errorf("Expected goblin to have 3 health, got %d", encounter[1].Health)
	}

	currentTurn = nextTurn(encounter, currentTurn)

	if currentTurn != 1 {
		t.Errorf("Expected current turn to be 1, got %d", currentTurn)
	}

	encounter, currentTurn = addEntities([]string{"add", "troll", "42", "25"}, encounter, currentTurn, true)

	// Troll becomes index 0, Alex is index 1, Goblin (active entity) should be index 2
	if currentTurn != 2 {
		t.Errorf("Expected current turn to be 2, got %d", currentTurn)
	}

	currentTurn = nextTurn(encounter, currentTurn)
	currentTurn = nextTurn(encounter, currentTurn)

	// Alex's turn
	if currentTurn != 1 {
		t.Errorf("Expected current turn to be 1, got %d", currentTurn)
	}

	fakeInputKill := bufio.NewScanner(strings.NewReader("y\n"))
	encounter, currentTurn = damageEntity([]string{"damage", "troll", "100"}, fakeInputKill, encounter, currentTurn)

	if len(encounter) != 2 {
		t.Errorf("Expected encounter to have length of 2, got %d", len(encounter))
	}

	if currentTurn != 0 {
		t.Errorf("Expected current turn to be 0, got %d", currentTurn)
	}
}
