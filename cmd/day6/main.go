package main

import (
	"aoc2023/internal/day6"
	"fmt"
)

func main() {
	result := day6.CalculatePart1([]day6.Race{
		{TimeMs: 38, Distance: 234},
		{TimeMs: 67, Distance: 1027},
		{TimeMs: 76, Distance: 1157},
		{TimeMs: 73, Distance: 1236},
	})
	fmt.Printf("Part1: %d\n", result)

	result2 := day6.CalculatePart1([]day6.Race{
		{TimeMs: 38677673, Distance: 234102711571236},
	})
	fmt.Printf("Part2: %d\n", result2)

}
