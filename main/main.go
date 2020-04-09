package main

import (
	"fmt"

	"github.com/AidanChueh/fantasy-basketball/packages/player"
	"github.com/AidanChueh/fantasy-basketball/packages/team"
)

func main() {
	steph := player.Create("Stephen Curry", 1)
	fmt.Println(steph)

	warriors := team.Create("Warriors")
	fmt.Println(warriors)
}
