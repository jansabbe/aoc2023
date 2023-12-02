package puzzle

import (
	"io"
	"strconv"
)

type Solver interface {
	Solve(reader io.Reader) (string, error)
}

type SolverFunc func(reader io.Reader) (string, error)

func (r SolverFunc) Solve(reader io.Reader) (string, error) {
	return r(reader)
}

type IntSolverFunc func(reader io.Reader) (int, error)

func (r IntSolverFunc) Solve(reader io.Reader) (string, error) {
	solve, err := r(reader)
	return strconv.Itoa(solve), err
}
