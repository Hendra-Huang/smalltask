package main

import (
	"container/heap"
	"fmt"
)

type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func insert(number int, minHeap *MinIntHeap, maxHeap *MaxIntHeap) {
	minHeapQueue := *minHeap
	if minHeap.Len() == 0 || number > minHeapQueue[0] {
		heap.Push(minHeap, number)
	} else {
		heap.Push(maxHeap, number)
	}
}

func rebalance(minHeap *MinIntHeap, maxHeap *MaxIntHeap) {
	diff := minHeap.Len() - maxHeap.Len()
	if diff >= 2 {
		x := heap.Pop(minHeap)
		heap.Push(maxHeap, x)
	} else if diff <= -2 {
		x := heap.Pop(maxHeap)
		heap.Push(minHeap, x)
	}
}

func getMedian(minHeap MinIntHeap, maxHeap MaxIntHeap) float32 {
	if minHeap.Len() == maxHeap.Len() {
		return float32(minHeap[0]+maxHeap[0]) * 0.5
	} else if minHeap.Len() > maxHeap.Len() {
		return float32(minHeap[0])
	} else {
		return float32(maxHeap[0])
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&numbers[i])
	}

	minHeap := new(MinIntHeap)
	maxHeap := new(MaxIntHeap)

	for _, number := range numbers {
		insert(number, minHeap, maxHeap)
		rebalance(minHeap, maxHeap)
		fmt.Printf("%0.1f\n", getMedian(*minHeap, *maxHeap))
	}
}
