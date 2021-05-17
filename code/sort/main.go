package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

var _ sort.IntSlice

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

func fullOrderBook() OrderHeap {
	var o OrderHeap
	heap.Push(&o, Order{Price: 1, Volume: 1})
	heap.Push(&o, Order{Price: 3, Volume: 3})
	heap.Push(&o, Order{Price: 4, Volume: 1})
	heap.Push(&o, Order{Price: 2, Volume: 2})
	heap.Push(&o, Order{Price: 6, Volume: 3})
	heap.Push(&o, Order{Price: 5, Volume: 2})
	return o
}

func printBook(book []Order) {
	var sl []string
	for _, o := range book {
		s := fmt.Sprintf("%d", o.Price)
		sl = append(sl, s)
	}
	fmt.Println(strings.Join(sl, ","))
}

// START OMIT
func main() {
	bookHeap := fullOrderBook()
	printBook(bookHeap)

	//arr := make([]Order, len(bookHeap))
	//copy(arr, bookHeap)
	//
	//sort.Slice(arr, func(i, j int) bool { return arr[i].Price < arr[j].Price })
	//
	//printBook(arr)
}

// END OMIT
