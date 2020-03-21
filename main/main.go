package main

import (
	"fmt"

	"github.com/AidanChueh/fantasy-basketball/packages/player"
)

func main() {
	steph := player.Create("Stephen Curry", "Point Guard")
	fmt.Println(steph)

	steph.UpdateName("Aidan Chueh")
	fmt.Println(steph)
}
