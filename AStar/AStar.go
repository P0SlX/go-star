package AStar

import (
	"github.com/P0SLX/go-star/node"
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

// AStar Implémentation de l'algorithme A*
func AStar(start, end *node.Node) []*node.Node {
	var openList, closedList []*node.Node
	var currentNode *node.Node

	openList = append(openList, start)

	for len(openList) > 0 {
		currentNode = openList[0]
		currentIndex := 0

		for index, v := range openList {
			if v.F < currentNode.F {
				currentNode = v
				currentIndex = index
			}
		}

		openList = append(openList[:currentIndex], openList[currentIndex+1:]...)
		closedList = append(closedList, currentNode)

		if currentNode == end {
			return reconstructPath(start, end)
		}

		for _, neighbor := range currentNode.Neighbors {
			if neighbor.IsWall || contains(closedList, neighbor) {
				continue
			}

			tempG := currentNode.G + 1

			if !contains(openList, neighbor) || tempG < neighbor.G {
				neighbor.G = tempG
				neighbor.H = Heuristic(neighbor, end)
				neighbor.F = neighbor.G + neighbor.H
				neighbor.Parent = currentNode

				if !contains(openList, neighbor) {
					openList = append(openList, neighbor)
				}
			}
		}
	}

	return nil
}
