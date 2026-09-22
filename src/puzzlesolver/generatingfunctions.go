package puzzlesolver

import (
	"math/rand"
	//"time"
)

func NewPuzzle(Size int) *Puzzle {
	newPuzzle := new(Puzzle)
	newPuzzle.Size = Size
	genPuzzle(newPuzzle)
	return newPuzzle
}

func genPuzzle(puzzle *Puzzle) {
	//rand.Seed(time.Now().UnixNano())
	for goodBoard := false; !goodBoard; goodBoard = checkPuzzle(*puzzle) {
		puzzle.Board = rand.Perm(puzzle.Size * puzzle.Size)
	}
}

func checkPuzzle(puzzle Puzzle) bool {
	inversions := 0

	if puzzle.Size%2 == 1 {
		for i := 1; i < len(puzzle.Board); i++ {
			seenThisInt := false
			for j := 1; j < len(puzzle.Board); j++ {
				if !seenThisInt && puzzle.Board[j] == i {
					seenThisInt = true
				} else if seenThisInt && puzzle.Board[j] < i {
					inversions++
				}
			}
		}
		return (inversions%2 == 0)
	}
	zeroIndex := -1
	for i := 0; i < len(puzzle.Board); i++ {
		if puzzle.Board[i] == 0 {
			zeroIndex = i
			break
		}
	}
	for i := 1; i < len(puzzle.Board); i++ {
		seenThisInt := false
		for j := 1; j < len(puzzle.Board); j++ {
			if !seenThisInt && puzzle.Board[j] == i {
				seenThisInt = true
			} else if seenThisInt && puzzle.Board[j] < i {
				inversions++
			}
		}
	}
	if (zeroIndex/puzzle.Size)%2 == 0 {
		return (inversions%2 == 0)
	} else {
		return (inversions%2 == 1)
	}
}

func CopyPuzzle(puz *Puzzle) *Puzzle {
	retPuz := new(Puzzle)
	retPuz.Size = puz.Size
	retPuz.Board = make([]int, 9)
	copy(puz.Board, retPuz.Board)
	return retPuz
}
