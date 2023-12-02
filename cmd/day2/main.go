package main

import (
	"aoc2023/internal/day2"
	"aoc2023/internal/puzzle"
)

func main() {
	puzzle.Solve(puzzle.Opts{
		Filename: "day2.txt",
		Part1:    puzzle.IntSolverFunc(day2.SumPossibleGames),
		Part2:    puzzle.IntSolverFunc(day2.CubeMinimumGame),
	})
}
