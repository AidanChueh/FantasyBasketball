package main

import (
	"fmt"

	"github.com/AidanChueh/fantasy-basketball/packages/player"
	"github.com/AidanChueh/fantasy-basketball/packages/team"
)

func main() {
	steph := player.Create("Stephen Curry", "Point Guard")
	fmt.Println(steph)

	fmt.Println(steph.Name())
	fmt.Println(steph.Position())

	steph.SetPosition("Center")
	fmt.Println(steph)

	warriors := team.Create("Warriors")
	fmt.Println(warriors.Players())

	warriors.AddPlayer(steph)
	fmt.Println(warriors)

	warriors.DeletePlayer("Stephen Curry")
	fmt.Println(warriors)
}
