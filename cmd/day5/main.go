package main

import (
	"aoc2023/internal/day5"
	"aoc2023/internal/puzzle"
)

func main() {
	puzzle.Solve(puzzle.Opts{
		Filename: "day5.txt",
		Part1:    puzzle.IntFunc(day5.CalculatePart1),
		Part2:    puzzle.IntFunc(day5.CalculatePart2),
	})
}
