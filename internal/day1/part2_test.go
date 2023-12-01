package day1

import (
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

func TestDay2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	actual, err := CalculatePart2(strings.NewReader(input))
	assert.NilError(t, err)
	assert.Equal(t, actual, 281)

}
