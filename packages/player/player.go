package player

import (
	"fmt"
)

// -----------
// Type Player
// -----------

// Player represents a basketball player
type Player struct {
	name string

	position position
}

// Create creates and returns a  player
func Create(name string, position position) Player {
	if position.isValid() {
		player := Player{name: name, position: position}
		return player
	}
	fmt.Println("Invalid position")
	player := Player{name: name, position: ""}
	return player
}

// Name returns the name
func (p *Player) Name() string {
	return (*p).name
}

// Position returns the position
func (p *Player) Position() position {
	return (*p).position
}

// SetPosition sets the position
func (p *Player) SetPosition(position position) {
	if position.isValid() {
		(*p).position = position
	} else {
		fmt.Println("Invalid position")
	}
}

// -------------
// Type Position
// -------------

type position string

const (
	pointGuard    position = "PG"
	shootingGuard position = "SG"
	smallForward  position = "SF"
	powerForward  position = "PF"
	center        position = "C"
)

func (p position) isValid() bool {
	if p == pointGuard || p == shootingGuard || p == smallForward || p == powerForward || p == center {
		return true
	}
	return false
}
