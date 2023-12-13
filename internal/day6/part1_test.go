package day6

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestParse(t *testing.T) {

	solution := CalculatePart1([]Race{
		{TimeMs: 7, Distance: 9},
		{TimeMs: 15, Distance: 40},
		{TimeMs: 30, Distance: 200},
	})

	assert.Equal(t, solution, 288)
}
