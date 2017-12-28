package main

import "fmt"

func main() {
	input := []int{5, 2, 7, 4, 3, 1, 6, 8, 9}
	fmt.Println(input)
	quicksort(input, 0, len(input)-1)
	fmt.Println(input)
}

func quicksort(numbers []int, start, end int) {
	if end <= start {
		return
	}
	partitionedIndex := partition(numbers, start, end)

	quicksort(numbers, start, partitionedIndex-1)
	quicksort(numbers, partitionedIndex+1, end)
}

func partition(numbers []int, start, end int) int {
	pivot := numbers[end]
	partitionedIndex := start - 1
	for i := start; i < end; i++ {
		if numbers[i] < pivot {
			partitionedIndex++
			numbers[partitionedIndex], numbers[i] = numbers[i], numbers[partitionedIndex]
		}
	}
	partitionedIndex++
	numbers[partitionedIndex], numbers[end] = numbers[end], numbers[partitionedIndex]

	return partitionedIndex
}
