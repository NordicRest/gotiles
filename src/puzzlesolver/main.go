package puzzlesolver

import "fmt"

//u "github.com/NordicRest/gotiles/src/utilities"
//s "github.com/davecgh/go-spew/spew"

// func testMovement() {
// 	npuz := u.NewPuzzle(3)

// 	f.Printf("Initial Puzzle: %v", npuz)
// 	wasgood, upuz := u.MoveFromTop(npuz)
// 	if wasgood {
// 		f.Printf("Moved from Top: %v\n", upuz)
// 	}
// 	wasgood, upuz = u.MoveFromBottom(npuz)
// 	if wasgood {
// 		f.Printf("Moved from Bottom: %v\n", upuz)
// 	}

// 	npuz = new(u.Puzzle)
// 	npuz.Size = 2
// 	npuz.Board = []int{1, 2, 3, 0}

// 	f.Printf("Puzzle %v is solved: %v\n", npuz, u.CheckSolved(npuz))

// 	npuz.Size = 3
// 	npuz.Board = []int{1, 2, 3, 4, 5, 6, 7, 8, 0}

// 	f.Printf("Puzzle %v is solved: %v\n", npuz, u.CheckSolved(npuz))

// 	npuz.Size = 2
// 	npuz.Board = []int{2, 1, 3, 0}

// 	f.Printf("Puzzle %v is solved: %v\n", npuz, u.CheckSolved(npuz))
// 	npuz.Size = 3
// 	npuz.Board = []int{2, 1, 3, 4, 5, 6, 7, 8, 0}

// 	f.Printf("Puzzle %v is solved: %v\n", npuz, u.CheckSolved(npuz))
// }

func Search(start *Puzzle, frontier SearchFrontier, heuristic Heuristic) *FrontierNode {
	startNode := &FrontierNode{
		Puz: start,
		Parent: nil,
		Depth: 0,
		CostToGo: heuristic(start),
	}
	
	bestCost := make(map[string]int)
	bestCost[PuzzleKey(start)] = 0

	frontier.Add(startNode)

	for frontier.Len() > 0 {
		current := frontier.GetNext()
		if CheckSolved(current.Puz) {
			return current
		}

		childDepth := current.Depth + 1

		tryMove := func(moveFunc func(*Puzzle) (bool, *Puzzle), move byte) {
			b, newChild := moveFunc(current.Puz)
			if !b {
				return
			}
			
			newKey := PuzzleKey(newChild)

			bestDepth, seen := bestCost[newKey]
			if seen && bestDepth <= childDepth {
				return
			}

			bestCost[newKey] = childDepth

			childNode := &FrontierNode{
				Puz:         newChild,
				Parent:      current,
				CurrentMove: move,
				Depth:       childDepth,
				CostToGo:    heuristic(newChild),
			}

			frontier.Add(childNode)
		}

		// Never
		tryMove(MoveFromTop, 't')

		// Eat
		tryMove(MoveFromTop, 'r')
		
		// Soggy
		tryMove(MoveFromTop, 'b')

		// Wieners
		tryMove(MoveFromTop, 'l')

	}

	return nil

}

func bfsHeuristic(p *Puzzle) int {
	return 0
}

func aStarHeuristic(p *Puzzle) int {
	return GetTaxicabDistance(p)
} 

func main() {
	testPuzzle := NewPuzzle(2)
	fmt.Printf("%+v\n", testPuzzle)
}
