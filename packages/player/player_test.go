package player

import (
	"testing"
)

// TO BE CREATED
// Compare with PlayerList
// func testValidName(t *testing.T, p Player) {}

func testValidPosition(t *testing.T, p Player) {
	validPositions := []string{"Point Guard", "Shooting Guard", "Small Forward", "Power Forward", "Center"}
	valid := false

	for _, validPosition := range validPositions {
		if validPosition == p.Position {
			valid = true
		}
	}

	if !valid {
		t.Errorf("Invalid position: %v", p.Position)
	}
}

func testValidUpdate(t *testing.T, new string, original string) {
	if new == original {
		t.Errorf("Expected updated value, but found same value: %v", new)
	}
}

func TestCreate(t *testing.T) {
	p := Create("Stephen Curry", "Point Guard")

	// testValidPlayer(t, p)
	testValidPosition(t, p)
}

func TestUpdateName(t *testing.T) {
	p := Create("Stephen Curry", "Point Guard")
	originalName := p.Name

	p.UpdateName("Aidan Chueh")
	newName := p.Name

	testValidUpdate(t, newName, originalName)
}

func TestUpdatePosition(t *testing.T) {
	p := Create("Stephen Curry", "Point Guard")
	originalPosition := p.Position

	p.UpdatePosition("Small Forward")
	newPosition := p.Position

	testValidUpdate(t, newPosition, originalPosition)
	testValidPosition(t, p)
}
