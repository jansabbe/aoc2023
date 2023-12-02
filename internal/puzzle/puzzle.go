package puzzle

import (
	"fmt"
	"log"
	"os"
)

type Opts struct {
	Filename string
	Part1    Solver
	Part2    Solver
}

func Solve(puzzle Opts) {
	file, err := os.Open(puzzle.Filename)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	if puzzle.Part1 == nil {
		return
	}

	part1, err := puzzle.Part1.Solve(file)
	if err != nil {
		log.Fatalf("Could not calculate part 1 because %v", err)
	}
	fmt.Printf("Part 1: %s\n", part1)

	if _, err := file.Seek(0, 0); err != nil {
		log.Fatalf("Could not seek to beginning of file, %v", err)
	}

	if puzzle.Part2 == nil {
		return
	}
	part2, err := puzzle.Part2.Solve(file)
	if err != nil {
		log.Fatalf("Could not calculate part 2 because %v", err)
	}
	fmt.Printf("Part 2: %s\n", part2)
}
