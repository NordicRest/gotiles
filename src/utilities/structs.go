package utilities

type FivePuzzle struct {
	board [5][5]uint8
}

type FourPuzzle struct {
	board [4][4]uint8
}

type ThreePuzzle struct {
	board [3][3]uint8
}

type Puzzle struct {
	size int
	board []int
}

