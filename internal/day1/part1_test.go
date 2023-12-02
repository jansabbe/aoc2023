package day1

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDay1(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	actual, err := CalculatePart1(strings.NewReader(input))
	assert.NilError(t, err)
	assert.Equal(t, actual, 142)
}
