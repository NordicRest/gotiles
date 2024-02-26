package main

import (
	f "fmt";
	u "github.com/NordicRest/gotiles/src/utilities"
	s "github.com/davecgh/go-spew/spew"
)

func testFrontierExpansion() {
	front := new(u.FrontierNode)
	f.Println("Initiated node")
	front.Puz = u.NewPuzzle(3)
	f.Println("initialized new puzzle")
	s.Dump(front.Puz)
	boole, val := u.ExpandNode(front)
	f.Println("expanded node")
	if !boole {s.Dump(val)}
	f.Println("Done")
}

func testMovement() {
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
	npuz.Size = 2
	npuz.Board = []int {1, 2, 3, 0}

	f.Println("Puzzle %v is solved: %b", npuz, u.CheckSolved(npuz))

	npuz.Size = 3
	npuz.Board = []int {1, 2, 3, 4, 5, 6, 7, 8, 0}

	f.Println("Puzzle %v is solved: %b", npuz, u.CheckSolved(npuz))

	npuz.Size = 2
	npuz.Board = []int {2, 1, 3, 0}

	f.Println("Puzzle %v is solved: %b", npuz, u.CheckSolved(npuz))
	npuz.Size = 3
	npuz.Board = []int {2, 1, 3, 4, 5, 6, 7, 8, 0}

	f.Println("Puzzle %v is solved: %b", npuz, u.CheckSolved(npuz))
}

func main() {
	
}