package playerlist

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/AidanChueh/fantasy-basketball/packages/player"
)

// PlayerList represents the entire list of players
type PlayerList []player.Player

// Create creates and returns a player list
func Create(filename string) (PlayerList, error) {
	playerList := PlayerList{}

	s := fileToSlice(filename)

	for i := range s {
		if i%2 != 0 {
			continue
		}

		position, err := strconv.Atoi(s[i+1])
		if err != nil {
			return nil, fmt.Errorf("invalid file contents. found %s at location %d", s[i+1], i+1)
		}
		p := player.Create(s[i], player.Position(position))
		playerList = append(playerList, p)
	}
	return playerList, nil
}

// Print prints the player list
func (p PlayerList) Print() {
	result := ""
	for i, player := range p {
		result += strconv.Itoa(i+1) + ". " + player.Name() + ", " + strconv.Itoa(int(player.Position())) + "\n"
	}
	fmt.Println(result)
}

// Random randomizes the player list
func (p PlayerList) Random() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range p {
		newPosition := r.Intn(len(p) - 1)
		p[i], p[newPosition] = p[newPosition], p[i]
	}
}

// Rank allows the user to ranks the players
func (p PlayerList) Rank() {
	var input string

	fmt.Println("Enter d when done ranking")

	for input != "d" {
		// Choosing player
		fmt.Println("\nEnter rank of player to move")
		fmt.Scanln(&input)
		if input != "d" {
			oldRank, _ := strconv.Atoi(input)
			player := p.getPlayer(oldRank)
			fmt.Println(player)

			// Moving player
			fmt.Println("Enter new rank")
			fmt.Scanln(&input)
			newRank, _ := strconv.Atoi(input)

			p.movePlayer(player, oldRank, newRank)
			p.Print()
		}
	}
	fmt.Println("Done")

}

// Choose allows the user to choose random or rank
func (p PlayerList) Choose() {
	var input string

	p.Print()

	fmt.Println("1. Rank")
	fmt.Println("2. Random")
	fmt.Println("\n Enter 1 or 2")
	fmt.Scanln(&input)
	if input == "1" {
		p.Rank()
	}
	if input == "2" {
		p.Random()
	}
}

// Contains determines whether the player list contains the player
func (p PlayerList) Contains(name string) bool {
	for _, player := range p {
		if player.Name() == name {
			return true
			// saves time instead of running through entire list
		}
	}
	return false
}

// Remove removes a player from the player list
func (p *PlayerList) Remove(name string) {
	for i, player := range *p {
		if player.Name() == name {
			*p = append((*p)[:i], (*p)[i+1:]...)
		}

	}
}

func fileToSlice(filename string) []string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// return error instead of println
		fmt.Println("Error:", err)
	}

	s := strings.Replace(string(bs), "\n", "", -1)
	return strings.Split(s, ",")
}

func (p PlayerList) getPlayer(rank int) player.Player {
	return p[rank-1]
}

func (p PlayerList) movePlayer(player player.Player, oldRank int, newRank int) {
	if newRank < oldRank {
		for i := oldRank - 2; i >= newRank-1; i-- {
			p[i+1] = p[i]
		}
		p[newRank-1] = player
	}

	for i := oldRank; i <= newRank-2; i++ {
		p[i-1] = p[i]
	}
	p[newRank-1] = player
}
