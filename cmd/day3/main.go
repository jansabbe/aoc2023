package main

import (
	"aoc2023/internal/day3"
	"aoc2023/internal/puzzle"
)

func main() {
	puzzle.Solve(puzzle.Opts{
		Filename: "day3.txt",
		Part1:    puzzle.IntFunc(day3.SumPartNumbers),
		Part2:    puzzle.IntFunc(day3.SumGears),
	})
}
