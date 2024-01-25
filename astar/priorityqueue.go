package astar

import (
	"container/heap"
	"github.com/P0SLX/go-star/node"
	"sort"
)

type PriorityQueue []*node.Node

// Check PriorityQueue implements sort.Interface
var _ sort.Interface = (*PriorityQueue)(nil)

// Check PriorityQueue implements heap.Interface
var _ heap.Interface = (*PriorityQueue)(nil)

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
