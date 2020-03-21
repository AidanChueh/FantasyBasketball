package player

// Player represents a basketball player
type Player struct {
	// Name represents the name of a basketball player
	Name string

	// Position represents the position of a basketball player
	Position string
}

// Create creates and returns a new player
func Create(newName string, newPosition string) Player {
	newPlayer := Player{Name: newName, Position: newPosition}
	return newPlayer
}

// UpdateName updates the name of the player
func (pointerToPlayer *Player) UpdateName(newName string) {
	(*pointerToPlayer).Name = newName
}

// UpdatePosition updates the position of the player
func (pointerToPlayer *Player) UpdatePosition(newPosition string) {
	(*pointerToPlayer).Position = newPosition
}
