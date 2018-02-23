package main

import (
	"fmt"
)

func getWordFrequency(words []string) map[string]int {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	return m
}

func main() {
	var magazineWordsCount, noteWordsCount int
	fmt.Scanln(&magazineWordsCount, &noteWordsCount)
	magazineWords := make([]string, magazineWordsCount)
	noteWords := make([]string, noteWordsCount)
	for i := 0; i < magazineWordsCount; i++ {
		fmt.Scan(&magazineWords[i])
	}
	for i := 0; i < noteWordsCount; i++ {
		fmt.Scan(&noteWords[i])
	}

	if noteWordsCount > magazineWordsCount {
		fmt.Println("No")
		return
	}

	magazineWordFrequency := getWordFrequency(magazineWords)
	noteWordFrequency := getWordFrequency(noteWords)

	for noteWord, noteFrequency := range noteWordFrequency {
		if magazineFrequency, ok := magazineWordFrequency[noteWord]; !ok || magazineFrequency < noteFrequency {
			fmt.Println("No")
			return
		}

		magazineWordFrequency[noteWord] -= noteFrequency
	}

	fmt.Println("Yes")
}
