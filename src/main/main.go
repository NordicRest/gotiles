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
}