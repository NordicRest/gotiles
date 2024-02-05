package utilities



func MoveFromTop(puz *Puzzle) (bool, *Puzzle) {
	zeroIndex := FindFirst(0, puz.board)
	if (zeroIndex < puz.size) {	// if the zero index is in the first row, we can't
		return false, nil		// move up, so return false/nil
	}

	retPuz := CopyPuzzle(puz)

	// swap 0 and the number "above" it
	retPuz.board[zeroIndex] = retPuz.board[zeroIndex - retPuz.size]
	retPuz.board[zeroIndex - retPuz.size] = 0

	return true, retPuz
}

func MoveFromBottom(puz *Puzzle) (bool, *Puzzle) {
	zeroIndex := FindFirst(0, puz.board)
	if (len(puz.board) - zeroIndex <= puz.size) {	// if the zero index is in the last row, we can't
		return false, nil							// move up, so return false/nil
	}

	retPuz := CopyPuzzle(puz)

	// swap 0 and the number "below" it
	retPuz.board[zeroIndex] = retPuz.board[zeroIndex + retPuz.size]
	retPuz.board[zeroIndex + retPuz.size] = 0

	return true, retPuz
}

func MoveFromLeft(puz *Puzzle) (bool, *Puzzle) {
	zeroIndex := FindFirst(0, puz.board)
	if (zeroIndex % puz.size == 0) {	// if the zero index is on the left side, we
		return false, nil				// can't move up, so return false/nil
	}

	retPuz := CopyPuzzle(puz)

	// swap 0 and the number before it
	retPuz.board[zeroIndex] = retPuz.board[zeroIndex - 1]
	retPuz.board[zeroIndex - 1] = 0

	return true, retPuz
}

func MoveFromRight(puz *Puzzle) (bool, *Puzzle) {
	zeroIndex := FindFirst(0, puz.board)
	if (zeroIndex % puz.size == (puz.size - 1)) {	// if the zero index is on the right side, we
		return false, nil							// can't move up, so return false/nil
	}

	retPuz := CopyPuzzle(puz)

	// swap 0 and the number "above" it
	retPuz.board[zeroIndex] = retPuz.board[zeroIndex - retPuz.size]
	retPuz.board[zeroIndex - retPuz.size] = 0

	return true, retPuz
}