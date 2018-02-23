package main

import (
	"fmt"
)

func getLetterFrequency(text string) map[rune]int {
	m := make(map[rune]int)
	for _, letter := range text {
		m[letter]++
	}

	return m
}

func sumLetterFrequency(letterFrequency map[rune]int) int {
	count := 0
	for _, frequency := range letterFrequency {
		count += frequency
	}

	return count
}

func main() {
	var first, second string
	fmt.Scanln(&first)
	fmt.Scanln(&second)
	firstLetterFrequency := getLetterFrequency(first)
	secondLetterFrequency := getLetterFrequency(second)

	var shorterLetterFrequency, longerLetterFrequency map[rune]int
	if len(firstLetterFrequency) < len(secondLetterFrequency) {
		shorterLetterFrequency = firstLetterFrequency
		longerLetterFrequency = secondLetterFrequency
	} else {
		shorterLetterFrequency = secondLetterFrequency
		longerLetterFrequency = firstLetterFrequency
	}

	for letter, firstFrequency := range shorterLetterFrequency {
		if secondFrequency, ok := longerLetterFrequency[letter]; ok {
			minimumFrequency := 0
			if firstFrequency < secondFrequency {
				minimumFrequency = firstFrequency
			} else {
				minimumFrequency = secondFrequency
			}
			shorterLetterFrequency[letter] -= minimumFrequency
			longerLetterFrequency[letter] -= minimumFrequency
		}
	}

	letterRemovedCount := sumLetterFrequency(shorterLetterFrequency) + sumLetterFrequency(longerLetterFrequency)
	fmt.Println(letterRemovedCount)
}
