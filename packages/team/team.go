package team

import (
	"github.com/AidanChueh/fantasy-basketball/packages/player"
)

// Team represents a team of players
type Team struct {
	name    string
	players []player.Player
}

// Create creates and returns a team
func Create(name string) Team {
	if name == "" {
		return Team{name: ""}
	}

	return Team{name: name}
}

// Players returns the players
func (t *Team) Players() []player.Player {
	return t.players
}

// AddPlayer adds a player to the team
func (t *Team) AddPlayer(player player.Player) {
	t.players = append(t.players, player)
}

// DeletePlayer deletes a player on the team
func (t *Team) DeletePlayer(name string) {
	for i, player := range t.players {
		if player.Name() == name {
			t.players = append(t.players[:i], (*t).players[i+1:]...)
		}
	}
}

func (t *Team) contains(name string) bool {
	contains := false

	for _, player := range t.players {
		if player.Name() == name {
			contains = true
		}
	}

	return contains
}
