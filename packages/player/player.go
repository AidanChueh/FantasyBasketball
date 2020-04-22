package player

// -----------
// Type Player
// -----------

// Player represents a basketball player
type Player struct {
	name string

	position Position
}

// Create creates and returns a player
func Create(name string, position Position) Player {
	if name == "" {
		return Player{name: "", position: position}
	}
	if !position.IsValid() {
		return Player{name: name, position: 0}
	}
	return Player{name: name, position: position}
}

// Name returns the name
func (p *Player) Name() string {
	return p.name
}

// Position returns the position
func (p *Player) Position() Position {
	return p.position
}

// SetPosition sets the position
func (p *Player) SetPosition(position Position) {
	p.position = position
}

// -------------
// Type Position
// -------------

// Position is the position of the basketball player
type Position int

// Possible position values
const (
	PointGuard Position = iota + 1
	ShootingGuard
	SmallForward
	PowerForward
	Center
)

// IsValid determines whether the position is valid
func (p Position) IsValid() bool {
	if p < PointGuard || p > Center {
		return false
	}

	return true
}

// Print array with numbers corresponding to positions
