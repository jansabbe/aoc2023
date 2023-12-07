package day5

import (
	"fmt"
	"io"
	"math"
)

func CalculatePart1(reader io.Reader) (int, error) {
	syllabus, err := Parse(reader)
	if err != nil {
		return 0, fmt.Errorf("could not parse syllabus: %w", err)
	}
	result := math.MaxInt
	for _, seed := range syllabus.seeds {
		calculated := syllabus.Calculate(seed)
		if result > calculated {
			result = calculated
		}
	}
	return result, nil
}
