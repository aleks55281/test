package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func calculateRating(goals, misses, assists int) float64 {
	if misses != 0 {
		return float64(((goals + assists) / 2) / misses)
	} else {
		return float64((goals + assists) / 2)
	}
}
func NewPlayer(name string, goals, misses, assists int) Player {

	rating := calculateRating(goals, misses, assists)
	return Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: rating}
}
func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Goals > players[j].Goals
	})
	return players
}
func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Rating > players[j].Rating
	})
	return players
}
func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		relI := players[i].Goals / players[i].Misses
		relJ := players[j].Goals / players[j].Misses
		if relI != relJ {
			return relI > relJ
		} else {
			return players[i].Name < players[j].Name
		}
	})
	return players
}

func main() {
	players := []Player{
		NewPlayer("Player1", 10, 5, 3),
		NewPlayer("Player2", 15, 7, 2),
		NewPlayer("Player3", 8, 2, 5),
	}
	fmt.Println(goalsSort(players))

}
