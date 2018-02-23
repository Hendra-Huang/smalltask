package main

import (
	"fmt"
)

type Operation struct {
	Type  string
	Value string
}

type TrieNode struct {
	Children map[rune]*TrieNode
	Value    string
	Total    int
}

func (n *TrieNode) Add(value string) {
	n.Total++
	currentNode := n
	valueLength := len(value)
	i := 0
	for ; i < valueLength; i++ {
		if node, ok := currentNode.Children[rune(value[i])]; ok {
			currentNode = node
			currentNode.Total++
		} else {
			break
		}
	}
	for ; i < valueLength; i++ {
		node := NewTrieNode()
		node.Total++
		currentNode.Children[rune(value[i])] = node
		currentNode = node
	}
	currentNode.Value = value
}

func (n *TrieNode) SearchCount(key string) int {
	currentNode := n
	for _, char := range key {
		if node, ok := currentNode.Children[char]; ok {
			currentNode = node
		} else {
			return 0
		}
	}

	return currentNode.Total
}

func NewTrieNode() *TrieNode {
	node := new(TrieNode)
	node.Children = make(map[rune]*TrieNode)

	return node
}

func main() {
	var operationCount int
	fmt.Scanln(&operationCount)
	operations := make([]Operation, operationCount)
	for i := 0; i < operationCount; i++ {
		var operationType, operationValue string
		fmt.Scanln(&operationType, &operationValue)
		operations[i] = Operation{
			Type:  operationType,
			Value: operationValue,
		}
	}

	root := NewTrieNode()
	for _, operation := range operations {
		if operation.Type == "add" {
			root.Add(operation.Value)
		} else if operation.Type == "find" {
			count := root.SearchCount(operation.Value)
			fmt.Println(count)
		}
	}
}
