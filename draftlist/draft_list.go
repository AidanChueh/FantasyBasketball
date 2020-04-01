package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/AidanChueh/fantasy-basketball/packages/player"
)

// DraftList represents a draft list of players
type DraftList []player.Player

// Random randomizes the draft list
func (d DraftList) Random() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// PlayerList represents the entire list of players
type PlayerList []player.Player

func fileToSlice(filename string) []string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)

	}
	s := strings.Split(string(bs), ",")
	return s
}

func newPlayerList(filename string) PlayerList {
	playerList := PlayerList{}

	// s := fileToSlice(filename)

	// for i, name := range s {
	// 	if i%2 == 0 {
	// 		p := player.Create(s[i], "")
	// 		p.SetPosition(s[i+1])
	// 		playerList = append(playerList, p)
	// 	}
	// }
	return playerList
}
