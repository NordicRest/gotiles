package utilities


func FindFirst(num int, board []int) int {
	for i := 0; i < len(board); i++ {
		if board[i] == num {
			return i
		}
	}
	return -1
}

func CheckSolved(puz *Puzzle) bool {
	for i := 0; i < len(puz.Board) - 1; i++ {
		if puz.Board[i] != i + 1 {
			return false
		}
	}
	return true
}