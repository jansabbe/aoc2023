package day2

import (
	"bufio"
	"io"
)

func SumImpossibleGame(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game, err := parseGame(line)
		if err != nil {
			return 0, err
		}
		if !game.IsImpossible(Set{Red: 12, Green: 13, Blue: 14}) {
			sum += game.Id
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}
