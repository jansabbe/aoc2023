package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func SumImpossibleGame(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game, err := parseGame(line)
		if err != nil {
			return 0, err
		}
		if !game.IsImpossible(Set{Red: 12, Green: 13, Blue: 14}) {
			sum += game.Id
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}

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

var redRegex = regexp.MustCompile("([0-9]+) red")
var greenRegex = regexp.MustCompile("([0-9]+) green")
var blueRegex = regexp.MustCompile("([0-9]+) blue")

func parseSet(value string) Set {
	return Set{
		Red:   getColorValue(value, redRegex),
		Blue:  getColorValue(value, blueRegex),
		Green: getColorValue(value, greenRegex),
	}
}

func getColorValue(value string, regex *regexp.Regexp) int {
	if matches := regex.FindStringSubmatch(value); matches != nil {
		result, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalf("matched %d but could not parse as int %v", matches[1], err)
		}
		return result
	}
	return 0
}

type Game struct {
	Id   int
	Sets []Set
}

func (g *Game) IsImpossible(limit Set) bool {
	for _, set := range g.Sets {
		if set.IsOverLimit(limit) {
			return true
		}
	}
	return false
}

type Set struct {
	Red, Green, Blue int
}

func (s *Set) IsOverLimit(limit Set) bool {
	return s.Red > limit.Red || s.Green > limit.Green || s.Blue > limit.Blue
}
