package main

import "fmt"

func main() {
	input := []int{5, 2, 7, 4, 3, 1, 6, 8, 9}
	sorted := quicksort(input)
	fmt.Println(input)
	fmt.Println(sorted)
}

func quicksort(numbers []int) []int {
	copiedNumbers := make([]int, len(numbers))
	copy(copiedNumbers, numbers)

	if len(copiedNumbers) <= 1 {
		return copiedNumbers
	}

	partitionedNumbers, partitionedIndex := partition(copiedNumbers)

	sortedNumbersStart := quicksort(partitionedNumbers[:partitionedIndex])
	sortedNumbersEnd := quicksort(partitionedNumbers[partitionedIndex+1:])

	sortedNumbersStart = append(sortedNumbersStart, partitionedNumbers[partitionedIndex])
	return append(sortedNumbersStart, sortedNumbersEnd...)
}

func partition(numbers []int) ([]int, int) {
	copiedNumbers := make([]int, len(numbers))
	copy(copiedNumbers, numbers)

	start := 0
	end := len(copiedNumbers) - 1

	pivot := copiedNumbers[end]
	partitionedIndex := start - 1
	for i := start; i < end; i++ {
		if copiedNumbers[i] < pivot {
			partitionedIndex++
			copiedNumbers[partitionedIndex], copiedNumbers[i] = copiedNumbers[i], copiedNumbers[partitionedIndex]
		}
	}
	partitionedIndex++
	copiedNumbers[partitionedIndex], copiedNumbers[end] = copiedNumbers[end], copiedNumbers[partitionedIndex]

	return copiedNumbers, partitionedIndex
}
