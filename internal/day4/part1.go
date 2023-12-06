package day4

import (
	"bufio"
	"io"
)

func CalculatePart1(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		card, err := ParseCard(line)
		if err != nil {
			return 0, err
		}
		sum += card.Score()
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}
