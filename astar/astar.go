package astar

import (
	"container/heap"
	"github.com/P0SLX/go-star/node"
	"github.com/P0SLX/go-star/utils"
)

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

// ColorPath Colorie le chemin trouvé en violet
func ColorPath(nodes []*node.Node) {
	for _, v := range nodes {
		v.Color.R = 255
		v.Color.G = 0
		v.Color.B = 255
	}
}

// AStar Algorithme A* unidirectionnel
func AStar(start, end *node.Node) []*node.Node {
	defer utils.Timer("astar")()
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

			tmpG := currentNode.G + 1
			if !neighbor.Already || neighbor.G > tmpG {
				neighbor.G = tmpG
				neighbor.H = neighbor.Heuristic(end)
				neighbor.F = neighbor.G + neighbor.H
				neighbor.Parent = currentNode
				neighbor.Already = true
				heap.Push(openSet, neighbor)
			}
		}
	}

	return nil
}
