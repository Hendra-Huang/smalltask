package main

import "fmt"

func main() {
	input := []int{5, 2, 7, 4, 3, 1, 6, 8, 9}
	fmt.Println(input)
	insertionsort(input, 0, len(input)-1)
	fmt.Println(input)
}

func insertionsort(numbers []int, start, end int) {
	for i := start + 1; i <= end; i++ {
		for j := i; j > start; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			} else {
				break
			}
		}
	}
}
