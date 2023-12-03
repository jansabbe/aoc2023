package day3

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	input := "467..114..\n" +
		"...*......\n" +
		"..35..633.\n" +
		"......#...\n" +
		"617*......\n" +
		".....+.58.\n" +
		"..592.....\n" +
		"......755.\n" +
		"...$.*....\n" +
		".664.598..\n"
	result, err := SumPartNumbers(strings.NewReader(input))
	assert.NilError(t, err)
	assert.Equal(t, result, 4361)
}
