package day3

import (
	"io"
)

func SumPartNumbers(input io.Reader) (int, error) {
	scanner := NewScanner(input)
	tokens := scanner.Tokens()

	sum := 0
	for _, token := range tokens {
		if token.Type == NUMBER {
			neighbours := token.FindNeighbours(tokens, SYMBOL)
			if len(neighbours) > 0 {
				sum += token.IntValue
			}
		}
	}
	return sum, nil
}
