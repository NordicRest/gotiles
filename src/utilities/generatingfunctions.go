package utilities

import (
	"math/rand"
	"fmt"
)

func NewThreePuzzle() *ThreePuzzle {
	newPuzzle := new(ThreePuzzle)
	genThreePuzzle(newPuzzle)
	return newPuzzle
}

func genThreePuzzle(puzzle *ThreePuzzle) {

	for goodBoard := false; !goodBoard; goodBoard = checkThreePuzzle(*puzzle) {

		arr := rand.Perm(9)
		fmt.Println("%d", arr)

		for i := 0; i < len(arr); i++ {
			*puzzle.board[i/3][i%3] = uint8(arr[i])
		}
		fmt.Println("%d", *puzzle.board)


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
	fmt.Println("%d", inversions)
	return (inversions % 2 == 0)
}
