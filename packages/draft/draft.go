package draft

import (
	"fmt"
	"strconv"

	"github.com/AidanChueh/fantasy-basketball/packages/player"

	"github.com/AidanChueh/fantasy-basketball/packages/playerlist"
	"github.com/AidanChueh/fantasy-basketball/packages/team"
)

// Draft represents a fantasy basketball draft
type Draft struct {
	globalList  playerlist.PlayerList
	playerList1 playerlist.PlayerList
	playerList2 playerlist.PlayerList
	team1       team.Team
	team2       team.Team
}

// Create creates a draft
func Create(filename string) Draft {
	globalList, _ := playerlist.Create(filename)
	playerList1, _ := playerlist.Create(filename)
	playerList2, _ := playerlist.Create(filename)
	team1 := team.Create("Team 1")
	team2 := team.Create("Team 2")

	return Draft{globalList: globalList, playerList1: playerList1, playerList2: playerList2, team1: team1, team2: team2}
}

// Ranking ranks the player lists
func (d Draft) Ranking() {
	fmt.Println("Team 1")
	d.playerList1.Choose()
	d.playerList1.Print()

	fmt.Println("Team 2")
	d.playerList2.Choose()
	d.playerList2.Print()
}

// Start starts the draft
func (d Draft) Start() {
	for i := 0; i <= len(d.globalList)-1; i++ {

		// Team 1's turn
		if i%2 == 0 {
			// Drafting a player to Team 1
			var target player.Player
			for _, player := range d.playerList1 {
				if d.globalList.Contains(player.Name()) {
					target = player
					d.team1.AddPlayer(target)
					fmt.Println("\nTeam 1 drafts " + target.Name())

					// Removing the player from global list
					d.globalList.Remove(target.Name())
					fmt.Println(d.globalList)
					break
				}
			}
		}

		// Team 2's turn
		if i%2 == 1 {
			// Drafting a player to Team 2
			var target player.Player
			for _, player := range d.playerList2 {
				if d.globalList.Contains(player.Name()) {
					target = player
					d.team2.AddPlayer(target)
					fmt.Println("\nTeam 2 drafts " + target.Name())

					// Removing the player from global list
					d.globalList.Remove(target.Name())
					break
				}
			}
		}
		fmt.Println("\nGlobal List length: " + strconv.Itoa(len(d.globalList)))
	}

	fmt.Println("Draft finished")
}

// PrintTeams prints the teams
func (d Draft) PrintTeams() {
	fmt.Println("Team 1:")
	fmt.Println(d.team1)
	fmt.Println("\nTeam 2: ")
	fmt.Println(d.team2)
}
