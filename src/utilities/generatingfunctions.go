package utilities

import (
	"math/rand"
	"time"
)

// func NewThreePuzzle() *ThreePuzzle {
// 	newPuzzle := new(ThreePuzzle)
// 	genThreePuzzle(newPuzzle)
// 	return newPuzzle
// }

// func genThreePuzzle(puzzle *ThreePuzzle) {

// 	rand.Seed(time.Now().UnixNano())
// 	for goodBoard := false; !goodBoard; goodBoard = checkThreePuzzle(*puzzle) {

// 		arr := rand.Perm(9)

// 		for i := 0; i < len(arr); i++ {
// 			puzzle.Board[i/3][i%3] = uint8(arr[i])
// 		}
// 	}
// }

// func checkThreePuzzle(puzzle ThreePuzzle) bool {
// 	inversions := 0
// 	for i := 1; i < len(puzzle.Board) * len(puzzle.Board); i++ {
// 		seenThisInt := false
// 		for j := 1; j < len(puzzle.Board) * len(puzzle.Board); j++ {
// 			if !seenThisInt && puzzle.Board[j/3][j%3] == uint8(i) {
// 				seenThisInt = true
// 			} else if seenThisInt && puzzle.Board[j/3][j%3] < uint8(i) {
// 				inversions++
// 			}
// 		}
// 	}
// 	return (inversions % 2 == 0)
// }

func NewPuzzle(Size int) *Puzzle {
	newPuzzle := new(Puzzle)
	newPuzzle.Size = Size
	genPuzzle(newPuzzle)
	return newPuzzle
}

func genPuzzle(puzzle *Puzzle) {
	rand.Seed(time.Now().UnixNano())
	for goodBoard := false; !goodBoard; goodBoard = checkPuzzle(*puzzle) {
		puzzle.Board = rand.Perm(puzzle.Size * puzzle.Size)
	}
}

func checkPuzzle(puzzle Puzzle) bool {
	inversions := 0

	if (puzzle.Size % 2 == 1){
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
	return (inversions % 2 == 0)
	}
	for i := 0; j < len(puzzle.Board); i++ {
		if puzzle.Board[i] == 0 {
			zeroIndex := i
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
	if (zeroIndex / puzzle.Size) % 2 == 0 {
		return (inversions % 2 == 0)
	} else {
		return (inversion % 2 == 1)
	}
}

func CopyPuzzle(puz *Puzzle) *Puzzle {
	retPuz := new(Puzzle)
	retPuz.Size = puz.Size
	retPuz.Board = make([]int, 9)
	for i := 0; i < len(puz.Board); i++ {
		retPuz.Board[i] = puz.Board[i]
	} 
	return retPuz
}