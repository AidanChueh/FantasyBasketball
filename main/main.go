package main

import "github.com/AidanChueh/fantasy-basketball/packages/draft"

func main() {
	draft := draft.Create("playerlist")
	draft.Start()

	draft.PrintTeams()
}
