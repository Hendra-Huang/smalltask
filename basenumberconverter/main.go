package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	ErrInvalidBase = errors.New("Invalid base")
	ValueToSymbol  = [20]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}
	SymbolToValue  = map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19}
)

func main() {
	var number string
	var fromBase, toBase int
	fmt.Print("Number: ")
	fmt.Scanf("%s", &number)
	fmt.Print("From Base: ")
	fmt.Scanf("%d", &fromBase)
	fmt.Print("To Base: ")
	fmt.Scanf("%d", &toBase)

	convertedNumber, err := convert(number, fromBase, toBase)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Result: %s\n", convertedNumber)
}

func convert(number string, fromBase, toBase int) (string, error) {
	if fromBase < 2 || fromBase > 20 || toBase < 2 || toBase > 20 {
		return "", ErrInvalidBase
	}

	if fromBase != 10 && toBase != 10 {
		decimal := convertToDecimal(number, fromBase)
		return convertFromDecimal(decimal, toBase), nil
	}
	if fromBase == 10 {
		if toBase == 10 {
			return number, nil
		}
		decimal, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return "", err
		}
		return convertFromDecimal(decimal, toBase), nil
	}
	if toBase == 10 {
		if fromBase == 10 {
			return number, nil
		}
		decimal := convertToDecimal(number, fromBase)
		return strconv.FormatInt(decimal, 10), nil
	}

	return "", nil
}

func convertFromDecimal(number int64, toBase int) string {
	var numberInString string
	for i := number; i > 0; i /= int64(toBase) {
		remainder := i % int64(toBase)
		numberInString = string(ValueToSymbol[remainder]) + numberInString
	}

	return numberInString
}

func convertToDecimal(number string, fromBase int) int64 {
	var decimal int64
	for i, j := 0, len(number)-1; j >= 0; i, j = i+1, j-1 {
		decimal += int64(SymbolToValue[rune(number[j])]) * int64(math.Pow(float64(fromBase), float64(i)))
	}

	return decimal
}
