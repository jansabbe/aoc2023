package day3

import (
	"io"
)

func SumGears(input io.Reader) (int, error) {
	scanner := NewScanner(input)
	tokens := scanner.Tokens()

	sum := 0
	for _, token := range tokens {
		if token.Type == SYMBOL && token.Value == "*" {
			neighbours := token.FindNeighbours(tokens, NUMBER)
			if len(neighbours) == 2 {
				sum += neighbours[0].IntValue * neighbours[1].IntValue
			}
		}
	}
	return sum, nil
}
