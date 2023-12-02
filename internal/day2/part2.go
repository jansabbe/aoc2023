package day2

import (
	"bufio"
	"io"
)

func CubeMinimumGame(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game, err := parseGame(line)
		if err != nil {
			return 0, err
		}
		sum += game.Power()
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}
