package puzzlesolver

import (
	"container/heap"
)

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

func NewAStarFrontier() *AStarFrontier {
	f := &AStarFrontier{}
	heap.Init(&f.pq)
	return f
}

func (front *AStarFrontier) Push(node *FrontierNode) {
	heap.Push(&front.pq, node)
}

func (front *AStarFrontier) Pop() *FrontierNode {
	if len(front.pq) == 0 {
		return nil
	}
	
	return heap.Pop(&front.pq).(*FrontierNode)
}

func (front *AStarFrontier) Len() int {
	return len(front.pq)
}

// internal priorityQueue actions abstracted by the above layer to make implementation consistent with BFS

func (pq *priorityQueue) Push(node any) {
	*pq = append(*pq, node.(*FrontierNode))
}

func (pq *priorityQueue) Pop() any {
	oldPQ := *pq
	n := len(oldPQ)

	node := oldPQ[n-1]
	oldPQ[n-1] = nil
	*pq  = oldPQ[:n-1]

	return node
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].NodeCost() < pq[j].NodeCost()
}