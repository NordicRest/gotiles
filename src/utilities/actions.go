package utilities

func MoveFromTop(puz *Puzzle) (bool, *Puzzle) {
	zeroIndex := FindFirst(0, puz.Board)
	if (zeroIndex < puz.Size) {	// if the zero index is in the first row, we can't
		return false, nil		// move up, so return false/nil
	}

	retPuz := CopyPuzzle(puz)

	// swap 0 and the number "above" it
	retPuz.Board[zeroIndex] = retPuz.Board[zeroIndex - retPuz.Size]
	retPuz.Board[zeroIndex - retPuz.Size] = 0

	return true, retPuz
}

func MoveFromBottom(puz *Puzzle) (bool, *Puzzle) {
	zeroIndex := FindFirst(0, puz.Board)
	if (len(puz.Board) - zeroIndex <= puz.Size) {	// if the zero index is in the last row, we can't
		return false, nil							// move up, so return false/nil
	}

	retPuz := CopyPuzzle(puz)

	// swap 0 and the number "below" it
	retPuz.Board[zeroIndex] = retPuz.Board[zeroIndex + retPuz.Size]
	retPuz.Board[zeroIndex + retPuz.Size] = 0

	return true, retPuz
}

func MoveFromLeft(puz *Puzzle) (bool, *Puzzle) {
	zeroIndex := FindFirst(0, puz.Board)
	if (zeroIndex % puz.Size == 0) {	// if the zero index is on the left side, we
		return false, nil				// can't move up, so return false/nil
	}

	retPuz := CopyPuzzle(puz)

	// swap 0 and the number before it
	retPuz.Board[zeroIndex] = retPuz.Board[zeroIndex - 1]
	retPuz.Board[zeroIndex - 1] = 0

	return true, retPuz
}

func MoveFromRight(puz *Puzzle) (bool, *Puzzle) {
	zeroIndex := FindFirst(0, puz.Board)
	if (zeroIndex % puz.Size == (puz.Size - 1)) {	// if the zero index is on the right side, we
		return false, nil							// can't move up, so return false/nil
	}

	retPuz := CopyPuzzle(puz)

	// swap 0 and the number "above" it
	retPuz.Board[zeroIndex] = retPuz.Board[zeroIndex - retPuz.Size]
	retPuz.Board[zeroIndex - retPuz.Size] = 0

	return true, retPuz
}