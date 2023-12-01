package day1

import (
	"bufio"
	"fmt"
	"io"
)

func CalculatePart2(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	numberToValues := createNumberToValueMap()
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		first, err := firstDigitOrNumber(line, numberToValues)
		if err != nil {
			return 0, err
		}

		last, err := lastDigitOrNumber(line, numberToValues)
		if err != nil {
			return 0, err
		}

		sum += (first * 10) + last
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}

func createNumberToValueMap() map[string]int {
	numbers := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := 0; i < 10; i++ {
		numbers[fmt.Sprintf("%d", i)] = i
	}
	return numbers
}

func firstDigitOrNumber(line string, numbers map[string]int) (int, error) {
	for i := 0; i < len(line); i++ {
		for digitOrNumber, value := range numbers {
			startIndex := i
			endIndex := i + len(digitOrNumber)

			if endIndex <= len(line) {
				actual := line[startIndex:endIndex]
				if actual == digitOrNumber {
					return value, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("no first digit or number found in line %q", line)

}

func lastDigitOrNumber(line string, numbers map[string]int) (int, error) {
	for i := len(line) - 1; i >= 0; i-- {
		for digitOrNumber, value := range numbers {
			startIndex := i - len(digitOrNumber) + 1
			endIndex := i + 1

			if startIndex >= 0 {
				actual := line[startIndex:endIndex]
				if actual == digitOrNumber {
					return value, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("no last digit or number found in line %q", line)
}
