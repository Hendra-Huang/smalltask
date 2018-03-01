package main

import "fmt"

var ErrStackEmpty = fmt.Errorf("Stack is empty")

type Stack struct {
	chars []rune
}

func (s *Stack) Push(char rune) {
	s.chars = append(s.chars, char)
}

func (s *Stack) Pop() (rune, error) {
	var char rune
	stackLength := s.Length()
	if stackLength == 0 {
		return char, ErrStackEmpty
	}

	char = s.chars[stackLength-1]
	s.chars = s.chars[:stackLength-1]

	return char, nil
}

func (s *Stack) Length() int {
	return len(s.chars)
}

func main() {
	var n int
	fmt.Scanln(&n)
	inputs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&inputs[i])
	}

	for _, input := range inputs {
		if isBalance(input) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isBalance(input string) bool {
	stack := new(Stack)
	for _, char := range input {
		if isOpenBracket(char) {
			stack.Push(char)
		} else if isCloseBracket(char) {
			stackChar, err := stack.Pop()
			if err != nil {
				return false
			}
			if !isMatchingPair(stackChar, char) {
				return false
			}
		} else {
			return false
		}
	}

	if stack.Length() > 0 {
		return false
	}

	return true
}

func isOpenBracket(char rune) bool {
	switch char {
	case '(':
		fallthrough
	case '[':
		fallthrough
	case '{':
		return true
	}

	return false
}

func isCloseBracket(char rune) bool {
	switch char {
	case ')':
		fallthrough
	case ']':
		fallthrough
	case '}':
		return true
	}

	return false
}

func isMatchingPair(open, close rune) bool {
	if (open == '(' && close == ')') ||
		(open == '[' && close == ']') ||
		(open == '{' && close == '}') {
		return true
	}

	return false
}
