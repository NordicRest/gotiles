package puzzlesolver

import (
	"fmt"
	"strings"
)

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

func PuzzleKey(p *Puzzle) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(p.Board)), ","), "[]")
}


func (n *FrontierNode) NodeCost() int {
	return n.Depth + n.CostToGo
}

func GetTaxicabDistance(n *Puzzle) int {
	distance := 0
	size := n.Size

	for i, currentVal := range n.Board {

		if currentVal == 0 {
			continue
		}

		goalIndex := currentVal - 1
		
		nDistance := (goalIndex % size) - (i % size)	// horizontal distance
		if nDistance < 0 {
			nDistance = -nDistance
		}
		distance += nDistance

		nDistance = (goalIndex / size) - (i / size)	// vertical distance
		if nDistance < 0 {
			nDistance = -nDistance
		}
		distance += nDistance
		
	}
	
	return distance
}