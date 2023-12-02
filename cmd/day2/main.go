package main

import (
	"aoc2023/internal/day2"
	"aoc2023/internal/puzzle"
)

func main() {
	puzzle.Solve(puzzle.Opts{
		Filename: "day2.txt",
		Part1:    puzzle.IntFunc(day2.SumPossibleGames),
		Part2:    puzzle.IntFunc(day2.CubeMinimumGame),
	})
}
