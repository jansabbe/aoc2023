package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func Parse(input io.Reader) (Syllabus, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(scanDoubleLines)

	if !scanner.Scan() {
		return Syllabus{}, fmt.Errorf("could not read input")
	}
	seeds := parseSeeds(scanner.Text())
	result := Syllabus{
		seeds: seeds,
	}
	for scanner.Scan() {
		rangeMap := ParseRangeMap(scanner.Text())
		result.rangeMaps = append(result.rangeMaps, rangeMap)
	}

	if err := scanner.Err(); err != nil {
		return Syllabus{}, err
	}

	return result, nil
}

func scanDoubleLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
