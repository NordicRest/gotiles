package main

import (
	f "fmt";
	u "github.com/NordicRest/gotiles/src/utilities"
)

func main() {
	npuz := u.NewPuzzle(3)
	f.Println("Initial Puzzle: %v", npuz)
	wasgood, upuz := u.MoveFromTop(npuz)
	if wasgood {
		f.Println("Moved from Top: %v", upuz)
	}
	wasgood, upuz = u.MoveFromBottom(npuz)
	if wasgood {
		f.Println("Moved from Bottom: %v", upuz)
	}

	npuz = new(u.Puzzle)
	npuz.size = 2
	npuz.board = []int {1, 2, 3, 0}

	f.Println("Puzzle %v is solved: %b", npuz, CheckSolved(npuz))
}