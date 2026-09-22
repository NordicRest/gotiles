package puzzlesolver

// type FrontierNode struct {
// 		Puz *Puzzle
// 		Parent *FrontierNode
// 		CurrentMove byte
// 		Depth int
// 		CostFromStart int
// 		CostToGo int
// }

// type SearchFrontier interface {
// 		Add(node *FrontierNode)
// 		GetNext() *FrontierNode
// 		Len() int
// }

// type BFSFrontier struct {
// 		nodes []*FrontierNode
// }

func (front *BFSFrontier) Push(node *FrontierNode) {
	front.nodes = append(front.nodes, node)
}

func (front *BFSFrontier) Pop() *FrontierNode {
	if len(front.nodes) == 0 {
		return nil
	}
	
	retNode := front.nodes[0]
	front.nodes[0] = nil
	front.nodes = front.nodes[1:]
	
	return retNode
}

func (front *BFSFrontier) Len() int {
	return len(front.nodes)
}