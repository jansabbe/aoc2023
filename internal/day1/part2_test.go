package day1

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDay2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	actual, err := CalculatePart2(strings.NewReader(input))
	assert.NilError(t, err)
	assert.Equal(t, actual, 281)

}
