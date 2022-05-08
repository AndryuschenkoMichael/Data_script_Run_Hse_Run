package model

import (
	"container/heap"
)

type edgeEntry struct {
	id   int
	cost float64
}

type Graph map[int][]edgeEntry
type distance map[int]float64

func (g Graph) AddEdge(edge Edge) {
	_, ok := g[edge.StartVertex]
	if !ok {
		g[edge.StartVertex] = []edgeEntry{}
	}

	g[edge.StartVertex] = append(g[edge.StartVertex], edgeEntry{
		id:   edge.EndVertex,
		cost: edge.Cost,
	})
}

func NewGraph(edges []Edge) Graph {
	graph := make(Graph)
	for _, edge := range edges {
		graph.AddEdge(edge)
	}

	return graph
}

func (g Graph) CountDistances(vertex Vertex) (edges []Edge) {
	edges = []Edge{}

	distances := make(distance)
	distances[vertex.Id] = 0
	heapEdge := newEdgeHeap()
	heap.Push(heapEdge, edgeEntry{
		id:   vertex.Id,
		cost: 0,
	})

	for heapEdge.Len() > 0 {
		v, _ := heap.Pop(heapEdge).(edgeEntry)
		dist, ok := distances[v.id]
		if !ok || dist < v.cost {
			continue
		}

		for _, value := range g[v.id] {
			d, ok := distances[value.id]
			if !ok || dist+value.cost < d {
				distances[value.id] = dist + value.cost
				heap.Push(heapEdge, edgeEntry{
					id:   value.id,
					cost: dist + value.cost,
				})
			}
		}
	}

	for i, value := range distances {
		edges = append(edges, Edge{
			StartVertex: vertex.Id,
			EndVertex:   i,
			Cost:        value,
			CampusId:    vertex.CampusId,
		})
	}
	return
}
