package team

import (
	"testing"

	"github.com/AidanChueh/fantasy-basketball/packages/player"
)

func testValidName(t *testing.T, team Team) {
	if team.name == "" {
		t.Errorf("Expected player name, but got %v", team.name)
	}
}

func TestCreate(t *testing.T) {
	team := Create("Warriors")

	testValidName(t, team)
}

func TestAddPlayer(t *testing.T) {
	team := Create("Warriors")
	player := player.Create("Stephen Curry", player.PointGuard)
	team.AddPlayer(player)

	if !team.contains("Stephen Curry") {
		t.Errorf("Expected team to include %v, but not found", player.Name())
	}
}

func TestDeletePlayer(t *testing.T) {
	team := Create("Warriors")
	player := player.Create("Stephen Curry", player.PointGuard)
	team.AddPlayer(player)
	team.DeletePlayer("Stephen Curry")

	if team.contains("Stephen Curry") {
		t.Errorf("Expected player to be deleted, but found %v", player.Name())
	}

}
