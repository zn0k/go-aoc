package paths

import (
	"container/heap"
	"math"

	"github.com/zn0k/go-aoc/queues"
)

func Dijkstra(xs []int, adjs AdjacencyList) map[int]int {
	distances := make(map[int]int)
	for i := 0; i < len(adjs); i += 1 {
		distances[i] = math.MaxInt32
	}
	for _, x := range xs {
		distances[x] = 0
	}

	visited := make(map[int]bool)

	q := make(queues.PriorityQueue, len(xs))
	for i := 0; i < len(xs); i += 1 {
		q[i] = &queues.Item{Value: xs[i], Priority: 0, Index: i}
	}
	heap.Init(&q)

	for q.Len() > 0 {
		item := heap.Pop(&q).(*queues.Item)
		_, done := visited[item.Value.(int)]
		if done {
			continue
		}
		visited[item.Value.(int)] = true

		for _, adj := range adjs[item.Value.(int)] {
			if distances[item.Value.(int)]+adj.weight < distances[adj.neighbor] {
				distances[adj.neighbor] = distances[item.Value.(int)] + adj.weight
				heap.Push(&q, &queues.Item{Value: adj.neighbor, Priority: distances[adj.neighbor]})
			}
		}
	}

	return distances
}
