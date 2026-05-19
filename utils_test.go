package main

import "testing"

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
