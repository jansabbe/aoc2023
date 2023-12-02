package main

import (
	"aoc2023/internal/day1"
	"aoc2023/internal/puzzle"
)

func main() {
	puzzle.Solve(puzzle.Opts{
		Filename: "day1.txt",
		Part1:    puzzle.IntSolverFunc(day1.CalculatePart1),
		Part2:    puzzle.IntSolverFunc(day1.CalculatePart2),
	})
}
