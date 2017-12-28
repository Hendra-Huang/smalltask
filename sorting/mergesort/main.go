package main

import "fmt"

func main() {
	input := []int{5, 2, 7, 4, 3, 1, 6, 8, 9}
	sorted := mergesort(input)
	fmt.Println(input)
	fmt.Println(sorted)
}

func mergesort(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	left := make([]int, 0)
	right := make([]int, 0)
	m := len(a) / 2

	for i, x := range a {
		switch {
		case i < m:
			left = append(left, x)
		case i >= m:
			right = append(right, x)
		}
	}

	left = mergesort(left)
	right = mergesort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {

	results := make([]int, 0)

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				results = append(results, left[0])
				left = left[1:len(left)]
			} else {
				results = append(results, right[0])
				right = right[1:len(right)]
			}
		} else if len(left) > 0 {
			results = append(results, left[0])
			left = left[1:len(left)]
		} else if len(right) > 0 {
			results = append(results, right[0])
			right = right[1:len(right)]
		}
	}

	return results
}
