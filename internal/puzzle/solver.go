package puzzle

import (
	"io"
	"strconv"
)

type Solver interface {
	Solve(reader io.Reader) (string, error)
}

type StringFunc func(reader io.Reader) (string, error)

func (r StringFunc) Solve(reader io.Reader) (string, error) {
	return r(reader)
}

type IntFunc func(reader io.Reader) (int, error)

func (r IntFunc) Solve(reader io.Reader) (string, error) {
	solve, err := r(reader)
	return strconv.Itoa(solve), err
}
