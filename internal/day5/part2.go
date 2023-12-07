package day5

import (
	"fmt"
	"io"
	"math"
)

func CalculatePart2(reader io.Reader) (int, error) {
	syllabus, err := Parse(reader)
	if err != nil {
		return 0, fmt.Errorf("could not parse syllabus: %w", err)
	}
	result := math.MaxInt

	for i := 0; i < len(syllabus.seeds); i += 2 {
		start := syllabus.seeds[i]
		length := syllabus.seeds[i+1]
		for seed := start; seed < start+length; seed++ {
			calculated := syllabus.Calculate(seed)
			if result > calculated {
				result = calculated
			}
		}
	}

	return result, nil
}
