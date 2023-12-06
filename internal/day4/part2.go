package day4

import (
	"bufio"
	"io"
)

func CalculatePart2(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	var cards []Card
	for scanner.Scan() {
		line := scanner.Text()
		card, err := ParseCard(line)
		if err != nil {
			return 0, err
		}
		cards = append(cards, card)
	}

	sum := 0
	for cardIdx, card := range cards {
		sum += card.Copies
		matchingNumbers := card.MatchingCards()
		for i := cardIdx + 1; i <= cardIdx+matchingNumbers; i++ {
			cards[i].Copies += card.Copies
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}
