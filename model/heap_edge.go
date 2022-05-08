package model

type edgeHeap []edgeEntry

func (h edgeHeap) Len() int {
	return len(h)
}

func (h edgeHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h edgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *edgeHeap) Push(x any) {
	*h = append(*h, x.(edgeEntry))
}

func (h *edgeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func newEdgeHeap() *edgeHeap {
	edgeHeap := edgeHeap{}
	return &edgeHeap
}