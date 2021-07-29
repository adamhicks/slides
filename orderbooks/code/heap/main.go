package main

import (
	"container/heap"
	"fmt"
)

type OrderHeap []Order

func (h OrderHeap) Len() int           { return len(h) }
func (h OrderHeap) Less(i, j int) bool { return h[i].Price < h[j].Price }
func (h OrderHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *OrderHeap) Push(x interface{}) {
	*h = append(*h, x.(Order))
}

func (h *OrderHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Order struct {
	Price  int
	Volume int
}

// START OMIT
func main() {
	var asks OrderHeap
	heap.Push(&asks, Order{Price: 10, Volume: 1})
	heap.Push(&asks, Order{Price: 9, Volume: 2})
	heap.Push(&asks, Order{Price: 11, Volume: 3})
	fmt.Println(heap.Pop(&asks))
}

// END OMIT
