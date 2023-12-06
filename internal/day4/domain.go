package day4

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	HaveNumbers    []int
	WinningNumbers []int
	Copies         int
}

func (c *Card) Score1() int {
	result := c.MatchingCards()
	if result == 0 {
		return 0
	}
	return int(math.Exp2(float64(result - 1)))
}

func (c *Card) MatchingCards() int {
	result := 0
	for _, winningNumber := range c.WinningNumbers {
		if slices.Contains(c.HaveNumbers, winningNumber) {
			result += 1
		}
	}
	return result
}

func ParseCard(line string) (Card, error) {
	_, numbers, found := strings.Cut(line, ":")
	if !found {
		return Card{}, fmt.Errorf("line does not match format %v", line)
	}
	winning, have, found := strings.Cut(numbers, "|")
	if !found {
		return Card{}, fmt.Errorf("line does not match format %v", line)
	}
	winningNumbers, err := parseNumbers(winning)
	if err != nil {
		return Card{}, fmt.Errorf("could not parse winning %v: %w", line, err)
	}
	haveNumbers, err := parseNumbers(have)
	if err != nil {
		return Card{}, fmt.Errorf("could not parse have %v: %w", line, err)
	}
	return Card{
		HaveNumbers:    haveNumbers,
		WinningNumbers: winningNumbers,
		Copies:         1,
	}, nil
}

func parseNumbers(numbersAsString string) ([]int, error) {
	splitNumbers := strings.Fields(numbersAsString)
	numbers := make([]int, len(splitNumbers))
	for i, value := range splitNumbers {
		valueAsInt, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("could not parse %v", value)
		}
		numbers[i] = valueAsInt
	}
	return numbers, nil
}
