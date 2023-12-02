package main

import (
	"aoc2023/internal/day2"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	part1, err := day2.SumImpossibleGame(file)
	if err != nil {
		log.Fatalf("Could not calculate part 1 because %v", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		log.Fatalf("Could not seek to beginning of file, %v", err)
	}

	fmt.Printf("Day 2, part 1: %d\n", part1)
}
