package main

import (
	"aoc2023/internal/day4"
	"aoc2023/internal/puzzle"
)

func main() {
	puzzle.Solve(puzzle.Opts{
		Filename: "day4.txt",
		Part1:    puzzle.IntFunc(day4.CalculatePart1),
	})
}
