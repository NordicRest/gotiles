package utilities


func FindFirst(num int, board []int) int {
	for i := 0; i < len(board); i++ {
		if board[i] == num {
			return i
		}
	}
	return -1
}