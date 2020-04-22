package player

import (
	"testing"
)

func testValidName(t *testing.T, p Player) {
	if p.Name() == "" {
		t.Errorf("Expected player name, but got %v", p.Name())
	}
}

func testValidPosition(t *testing.T, p Player) {
	if p.Position() < PointGuard || p.Position() > Center {
		t.Errorf("Expected position of 1-5, but got %v", p.Position())
	}
}

func TestCreate(t *testing.T) {
	p := Create("Stephen Curry", 19)

	testValidName(t, p)
	testValidPosition(t, p)

}

func TestPositionSetPosition(t *testing.T) {
	p := Create("Player", 3)

	position := SmallForward
	p.SetPosition(position)

	if p.Position() != position {
		t.Errorf("Expected position of %v, but got %v", position, p.Position())
	}
}
