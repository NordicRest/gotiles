package utilities

import (
	"math/rand"
	"time"
)

func NewThreePuzzle() *ThreePuzzle {
	newPuzzle := new(ThreePuzzle)
	genThreePuzzle(newPuzzle)
	return newPuzzle
}

func genThreePuzzle(puzzle *ThreePuzzle) {

	rand.Seed(time.Now().UnixNano())
	for goodBoard := false; !goodBoard; goodBoard = checkThreePuzzle(*puzzle) {

		arr := rand.Perm(9)

		for i := 0; i < len(arr); i++ {
			puzzle.board[i/3][i%3] = uint8(arr[i])
		}
	}
}

func checkThreePuzzle(puzzle ThreePuzzle) bool {
	inversions := 0
	for i := 1; i < len(puzzle.board) * len(puzzle.board); i++ {
		seenThisInt := false
		for j := 1; j < len(puzzle.board) * len(puzzle.board); j++ {
			if !seenThisInt && puzzle.board[j/3][j%3] == uint8(i) {
				seenThisInt = true
			} else if seenThisInt && puzzle.board[j/3][j%3] < uint8(i) {
				inversions++
			}
		}
	}
	return (inversions % 2 == 0)
}

func NewPuzzle(size int) *Puzzle {
	newPuzzle := new(Puzzle)
	newPuzzle.size = size
	genPuzzle(newPuzzle)
	return newPuzzle
}

func genPuzzle(puzzle *Puzzle) {
	rand.Seed(time.Now().UnixNano())
	for goodBoard := false; !goodBoard; goodBoard = checkPuzzle(*puzzle) {
		puzzle.board = rand.Perm(puzzle.size * puzzle.size)
	}
}

func checkPuzzle(puzzle Puzzle) bool {
	inversions := 0

	if (puzzle.size % 2 == 1){
		for i := 1; i < len(puzzle.board); i++ {
			seenThisInt := false
			for j := 1; j < len(puzzle.board); j++ {
				if !seenThisInt && puzzle.board[j] == i {
					seenThisInt = true
				} else if seenThisInt && puzzle.board[j] < i {
					inversions++
				}
			}
		}
	return (inversions % 2 == 0)
	}
	return false
}

func CopyPuzzle(puz *Puzzle) *Puzzle {
	retPuz := new(Puzzle)
	retPuz.size = puz.size
	retPuz.board = make([]int, 9)
	for i := 0; i < len(puz.board); i++ {
		retPuz.board[i] = puz.board[i]
	} 
	return retPuz
}