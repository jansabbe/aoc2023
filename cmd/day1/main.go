package main

import (
	"aoc2023/internal/day1"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer file.Close()

	part1, err := day1.CalculatePart1(file)
	if err != nil {
		log.Fatal("Could not calculate part 1", err)
	}

	file.Seek(0, 0)
	part2, err := day1.CalculatePart2(file)
	if err != nil {
		log.Fatal("Could not calculate part 2", err)
	}

	fmt.Printf("Day 1, part 1: %d\n", part1)
	fmt.Printf("Day 1, part 2: %d\n", part2)
}
