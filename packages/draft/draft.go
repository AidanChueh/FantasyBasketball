package draft

import (
	"fmt"
	"strconv"

	"github.com/AidanChueh/fantasy-basketball/packages/playerlist"
	"github.com/AidanChueh/fantasy-basketball/packages/team"
)

// Draft represents a fantasy basketball draft
type Draft struct {
	globalList playerlist.PlayerList
	playerList []playerlist.PlayerList
	team       []team.Team
}

// Create creates a draft
func Create(filename string) Draft {
	globalList, _ := playerlist.Create(filename)
	playerList1, _ := playerlist.Create(filename)
	playerList2, _ := playerlist.Create(filename)

	return CreateArgs(globalList, playerList1, playerList2)
}

// CreateArgs creates a draft with any number of lists and teams
func CreateArgs(g playerlist.PlayerList, draftPreferences ...playerlist.PlayerList) Draft {
	d := Draft{globalList: g, playerList: []playerlist.PlayerList{}, team: []team.Team{}}

	for i, list := range draftPreferences {
		d.playerList = append(d.playerList, list)
		d.team = append(d.team, team.Create("Team "+strconv.Itoa(i)))
	}

	return d
}

// Ranking ranks the player lists
func (d Draft) Ranking() {

	for i, list := range d.playerList {
		list.Choose()

		fmt.Println("Team " + strconv.Itoa(i))
		list.Print()
	}
}

// Start starts the draft
func (d *Draft) Start() {
	for range d.globalList {

		for j, list := range d.playerList {

			for _, player := range list {
				if d.globalList.Contains(player.Name()) {

					d.team[j].AddPlayer(player)
					fmt.Printf("Team %d drafts %s\n", j, player.Name())

					// Removing the player from global list
					d.globalList.Remove(player.Name())
					break
				}
			}
		}
	}

	fmt.Println("Draft finished")
}

// PrintTeams prints the teams
func (d Draft) PrintTeams() {
	for i, team := range d.team {
		fmt.Printf("Team %d:", i)
		fmt.Println(team)
	}
}
