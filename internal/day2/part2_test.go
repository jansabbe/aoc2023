package day2

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart2(t *testing.T) {
	input := strings.NewReader("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n" +
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n" +
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n" +
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n" +
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")
	result, err := CubeMinimumGame(input)
	assert.NilError(t, err)
	assert.Equal(t, result, 2286)
}
