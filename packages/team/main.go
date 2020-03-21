package team

import (
	"fmt"

	"github.com/AidanChueh/fantasy-basketball/packages/player"
)

func main() {
	steph := player.Create("Stephen Curry", "Point Guard")
	klay := player.Create("Klay Thompson", "Shooting Guard")
	kd := player.Create("Kevin Durant", "Small Forward")

	// Testing Create
	warriors := Create("Warriors")

	// Testing AddPlayer
	warriors.AddPlayer(steph)
	warriors.AddPlayer(klay)
	warriors.AddPlayer(kd)
	fmt.Println(warriors)

	// Testing DeletePlayer
	warriors.DeletePlayer("Klay Thompson")
	fmt.Println(warriors)
}
