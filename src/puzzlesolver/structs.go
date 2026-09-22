package puzzlesolver

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
	Size int
	Board []int
}

type FrontierNode struct {
	Puz *Puzzle
	Parent *FrontierNode
	CurrentMove byte
	Depth int
	CostToGo int
}

type SearchFrontier interface {
	Add(node *FrontierNode)
	GetNext() *FrontierNode
	Len() int
}

type BFSFrontier struct {
	nodes []*FrontierNode
}

type priorityQueue [] *FrontierNode

type AStarFrontier struct {
	pq priorityQueue
}

type Heuristic func(*Puzzle) int 