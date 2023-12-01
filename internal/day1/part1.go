package day1

import (
	"bufio"
	"fmt"
	"io"
)

func CalculatePart1(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		first, err := firstDigit(line)
		if err != nil {
			return 0, err
		}

		last, err := lastDigit(line)
		if err != nil {
			return 0, err
		}
		sum += (first * 10) + last
	}
	return sum, nil
}

func firstDigit(line string) (int, error) {
	for i := 0; i < len(line); i++ {
		char := line[i]
		if char > '0' && char <= '9' {
			return int(char - '0'), nil
		}
	}
	return 0, fmt.Errorf("no first digit found in line [%s]", line)
}

func lastDigit(line string) (int, error) {
	for i := len(line) - 1; i >= 0; i-- {
		char := line[i]
		if char > '0' && char <= '9' {
			return int(char - '0'), nil
		}
	}
	return 0, fmt.Errorf("no last digit found in line [%s]", line)
}
