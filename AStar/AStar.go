package AStar

import (
	"container/heap"
	"github.com/P0SLX/go-star/node"
	"github.com/P0SLX/go-star/utils"
	"math"
)

// Heuristic calcule la distance Euclidienne entre 2 points
func Heuristic(node, dest *node.Node) float64 {
	xSquare := float64(node.X-dest.X) * float64(node.X-dest.X)
	ySquare := float64(node.Y-dest.Y) * float64(node.Y-dest.Y)
	return math.Sqrt(xSquare + ySquare)
}

func reconstructPath(start, end *node.Node) []*node.Node {
	var path []*node.Node
	currentNode := end

	for currentNode != start {
		path = append(path, currentNode)
		currentNode = currentNode.Parent
	}

	path = append(path, start)

	return path
}

func contains(nodes []*node.Node, node *node.Node) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}

	return false
}

// ColorPath Colorie le chemin trouvé en violet
func ColorPath(nodes []*node.Node) {
	for _, v := range nodes {
		v.Color.R = 255
		v.Color.G = 0
		v.Color.B = 255
	}
}

type PriorityQueue []*node.Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*node.Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func AStar(start, end *node.Node) []*node.Node {
	defer utils.Timer("AStar")()
	openSet := &PriorityQueue{}
	closedSet := make(map[*node.Node]bool)
	heap.Init(openSet)
	heap.Push(openSet, start)

	for openSet.Len() > 0 {
		currentNode := heap.Pop(openSet).(*node.Node)

		if currentNode == end {
			return reconstructPath(start, end)
		}

		closedSet[currentNode] = true

		for _, neighbor := range currentNode.Neighbors {
			if closedSet[neighbor] || neighbor.IsWall {
				continue
			}

			if !contains(*openSet, neighbor) {
				neighbor.G = currentNode.G + 1
				neighbor.H = Heuristic(neighbor, end)
				neighbor.F = neighbor.G + neighbor.H
				neighbor.Parent = currentNode
				heap.Push(openSet, neighbor)
			}
		}
	}

	return nil
}
