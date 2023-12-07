package day5

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Seeds []int

func parseSeeds(line string) Seeds {
	fields := strings.Fields(line)[1:]
	return mapToInts(fields)
}

func mapToInts(fields []string) []int {
	result := make([]int, len(fields))
	for i, val := range fields {
		converted, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		result[i] = converted
	}
	return result
}

type Syllabus struct {
	seeds     Seeds
	rangeMaps []RangeMap
}

func (r *Syllabus) Calculate(seed int) int {
	result := seed
	for _, rangeMap := range r.rangeMaps {
		idx, found := slices.BinarySearchFunc(rangeMap, result, func(r Range, i int) int {
			if r.SourceFrom <= i && i <= r.SourceTo {
				return 0
			}
			return cmp.Compare(r.SourceFrom, i)
		})
		if found {
			result += rangeMap[idx].Delta
		}
	}
	return result
}

func (r *Syllabus) SeedRange() []int {
	var result []int
	for i := 0; i < len(r.seeds); i += 2 {
		start := r.seeds[i]
		length := r.seeds[i+1]
		for j := start; j < start+length; j++ {
			result = append(result, j)
		}
	}
	return result
}

type RangeMap []Range

type Range struct {
	SourceFrom int
	SourceTo   int
	Delta      int
}

func ParseRangeMap(rangeMapAsString string) RangeMap {
	lines := strings.Split(rangeMapAsString, "\n")[1:]
	rangeMap := make([]Range, len(lines))
	for i, line := range lines {
		nums := mapToInts(strings.Fields(line))
		rangeMap[i] = Range{
			SourceFrom: nums[1],
			SourceTo:   nums[1] + nums[2] - 1,
			Delta:      -nums[1] + nums[0],
		}
	}
	slices.SortFunc(rangeMap, func(a, b Range) int {
		return cmp.Compare(a.SourceFrom, b.SourceFrom)
	})
	return rangeMap
}
