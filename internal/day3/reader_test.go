package day3

import (
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

func TestScanner_Tokens(t *testing.T) {
	input := strings.NewReader("467..114..\n" +
		"...*......\n" +
		"..35..633.\n")
	scanner := NewScanner(input)
	tokens := scanner.Tokens()
	assert.DeepEqual(t, tokens, []Token{
		{
			Type:       NUMBER,
			Line:       0,
			StartIndex: 0,
			EndIndex:   2,
			Value:      "",
			IntValue:   467,
		},
		{
			Type:       NUMBER,
			Line:       0,
			StartIndex: 5,
			EndIndex:   7,
			Value:      "",
			IntValue:   114,
		},
		{
			Type:       SYMBOL,
			Line:       1,
			StartIndex: 3,
			EndIndex:   3,
			Value:      "*",
			IntValue:   0,
		},
		{
			Type:       NUMBER,
			Line:       2,
			StartIndex: 2,
			EndIndex:   3,
			Value:      "",
			IntValue:   35,
		},
		{
			Type:       NUMBER,
			Line:       2,
			StartIndex: 6,
			EndIndex:   8,
			Value:      "",
			IntValue:   633,
		},
	})
}
