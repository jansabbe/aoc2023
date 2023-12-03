package day3

import (
	"io"
	"slices"
)

func SumPartNumbers(input io.Reader) (int, error) {
	scanner := NewScanner(input)
	tokens := scanner.Tokens()

	sum := 0
	for _, token := range tokens {
		if token.Type == NUMBER {
			neighbours := findNeighbours(tokens, token)
			hasSymbolAsNeighbour := slices.ContainsFunc(neighbours, func(token Token) bool {
				return token.Type == SYMBOL
			})
			if hasSymbolAsNeighbour {
				sum += token.IntValue
			}
		}
	}
	return sum, nil
}

func findNeighbours(tokens []Token, token Token) []Token {
	var result []Token
	for _, potentialNeighbour := range tokens {
		if token.IsNeighbour(potentialNeighbour) {
			result = append(result, potentialNeighbour)
		}
	}
	return result
}
