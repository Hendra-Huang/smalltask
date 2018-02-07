/*
Given a list of transactions, How can we calculate the frequency counts of all possible item-sets?
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func generateCombination(choices []string, choose int) [][]string {
	choicesLength := len(choices)
	combinations := [][]string{}
	combination := make([]string, choose)

	var recursiveWalk func(i, next int)
	recursiveWalk = func(i, next int) {
		for choiceID := next; choiceID < choicesLength; choiceID++ {
			combination[i] = choices[choiceID]
			if i < (choose - 1) {
				recursiveWalk(i+1, choiceID+1)
			} else {
				copied := make([]string, choose)
				copy(copied, combination)
				combinations = append(combinations, copied)
			}
		}
	}
	recursiveWalk(0, 0)

	return combinations
}

func main() {
	choicesInput := [][]string{
		[]string{"apple", "lemon", "banana"},
		[]string{"banana", "berry", "lemon", "orange"},
		[]string{"banana", "berry", "lemon"},
	}

	longestChoicesLength := 0
	for _, val := range choicesInput {
		length := len(val)
		if length > longestChoicesLength {
			longestChoicesLength = length
		}
	}
	results := make(map[string]int)
	for choose := 2; choose <= longestChoicesLength; choose++ {
		for _, val := range choicesInput {
			combinationPairs := generateCombination(val, choose)
			for _, combinationPair := range combinationPairs {
				sort.Strings(combinationPair)
				key := strings.Join(combinationPair, ",")
				results[key]++
			}
		}
	}

	fmt.Printf("%+v\n", results)
}
