package day2

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parseGame(line string) (Game, error) {
	before, after, found := strings.Cut(line, ":")
	if !found {
		return Game{}, fmt.Errorf("no : found in %v", line)
	}

	var id int
	_, err := fmt.Sscanf(before, "Game %d", &id)
	if err != nil {
		return Game{}, fmt.Errorf("expected line to start with game %v: %w", line, err)
	}
	sets := parseSets(after)

	return Game{
		Id:   id,
		Sets: sets,
	}, nil
}

func parseSets(input string) []Set {
	setsAsStrings := strings.Split(input, ";")
	sets := make([]Set, len(setsAsStrings))
	for i, value := range setsAsStrings {
		sets[i] = parseSet(value)
	}
	return sets
}

var (
	redRegex   = regexp.MustCompile("([0-9]+) red")
	greenRegex = regexp.MustCompile("([0-9]+) green")
	blueRegex  = regexp.MustCompile("([0-9]+) blue")
)

func parseSet(value string) Set {
	return Set{
		Red:   getColorValue(value, redRegex),
		Blue:  getColorValue(value, blueRegex),
		Green: getColorValue(value, greenRegex),
	}
}

func getColorValue(value string, regex *regexp.Regexp) int {
	matches := regex.FindStringSubmatch(value)
	if matches == nil {
		return 0
	}
	result, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatalf("matched %v but could not parse as int %v", matches[1], err)
	}
	return result
}
