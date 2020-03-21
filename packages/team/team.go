package team

import (
	"github.com/AidanChueh/fantasy-basketball/packages/player"
)

// Team represents a team of players
type Team struct {
	name    string
	players []player.Player
}

// Create creates and returns a new team
func Create(newName string) Team {
	newTeam := Team{name: newName}
	return newTeam
}

// UpdateName updates the name of the team
func (pointerToTeam *Team) UpdateName(newName string) {
	(*pointerToTeam).name = newName
}

// AddPlayer adds a player to the team
func (pointerToTeam *Team) AddPlayer(newPlayer player.Player) {
	(*pointerToTeam).players = append((*pointerToTeam).players, newPlayer)
}

// DeletePlayer deletes a player on the team
func (pointerToTeam *Team) DeletePlayer(name string) {
	for i, player := range (*pointerToTeam).players {
		if name == player.Name {
			(*pointerToTeam).players = append((*pointerToTeam).players[:i], (*pointerToTeam).players[i+1:]...)
		}
	}
}
