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

// AStarBiDirectional Algorithme A* bidirectionnel (non fonctionnel)
func AStarBiDirectional(start, end *node.Node) []*node.Node {
	defer utils.Timer("astar")()

	// Initialisation des files de priorité pour les deux directions
	openSetStart, openSetEnd := &PriorityQueue{}, &PriorityQueue{}
	closedSetStart, closedSetEnd := make(map[*node.Node]bool), make(map[*node.Node]bool)

	heap.Init(openSetStart)
	heap.Init(openSetEnd)
	heap.Push(openSetStart, start)
	heap.Push(openSetEnd, end)

	// Channel pour communiquer la rencontre
	meetPoint := make(chan *node.Node)

	var pathStart, pathEnd []*node.Node

	go searchPath(openSetStart, closedSetStart, end, meetPoint)
	go searchPath(openSetEnd, closedSetEnd, start, meetPoint)

	// Attendre que les deux recherches se rencontrent
	meetNode := <-meetPoint

	// Reconstruire le chemin depuis le point de rencontre
	if meetNode != nil {
		pathStart = reconstructPath(start, meetNode)
		pathEnd = reconstructPath(end, meetNode)
		return mergePaths(pathStart, pathEnd)
	} else {
		return reconstructPath(start, end)
	}
}

// La fonction searchPath est similaire à la fonction AStar initiale mais prend en compte la direction
// et communique via le channel meetPoint lorsqu'un nœud commun est trouvé.
func searchPath(openSet *PriorityQueue, closedSet map[*node.Node]bool, target *node.Node, meetPoint chan<- *node.Node) {
	for openSet.Len() > 0 {
		currentNode := heap.Pop(openSet).(*node.Node)

		if closedSet[currentNode] {
			meetPoint <- currentNode
			return
		}

		closedSet[currentNode] = true

		for _, neighbor := range currentNode.Neighbors {
			if closedSet[neighbor] || neighbor.IsWall {
				continue
			}

			tmpG := currentNode.G + 1
			if !neighbor.Already || neighbor.G > tmpG {
				neighbor.G = tmpG
				neighbor.H = neighbor.Heuristic(target)
				neighbor.F = neighbor.G + neighbor.H
				neighbor.Parent = currentNode
				neighbor.Already = true
				heap.Push(openSet, neighbor)

				// Colorier les nœuds explorés
				neighbor.Color.R = uint8(neighbor.G * 10)
				neighbor.Color.G = 100
				neighbor.Color.B = 255
			}
		}
	}

	meetPoint <- nil // Aucun point de rencontre trouvé
}

// La fonction mergePaths fusionne les deux chemins obtenus
func mergePaths(pathStart, pathEnd []*node.Node) []*node.Node {
	// Inverser le chemin de fin pour qu'il parte du point de rencontre
	reverse(pathEnd)

	// Fusionner les deux chemins
	return append(pathStart, pathEnd...)
}

func reverse(nodes []*node.Node) {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
}
