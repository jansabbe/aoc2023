package day2

import (
	"cmp"
	"slices"
)

type Game struct {
	Id   int
	Sets []Set
}

func (g *Game) IsPossible(limit Set) bool {
	return !slices.ContainsFunc(g.Sets, func(set Set) bool {
		return set.IsOverLimit(limit)
	})
}

func (g *Game) Power() int {
	maximumRed := slices.MaxFunc(g.Sets, func(a, b Set) int {
		return cmp.Compare(a.Red, b.Red)
	})
	maximumGreen := slices.MaxFunc(g.Sets, func(a, b Set) int {
		return cmp.Compare(a.Green, b.Green)
	})
	maximumBlue := slices.MaxFunc(g.Sets, func(a, b Set) int {
		return cmp.Compare(a.Blue, b.Blue)
	})
	return maximumRed.Red * maximumGreen.Green * maximumBlue.Blue
}

type Set struct {
	Red, Green, Blue int
}

func (s *Set) IsOverLimit(limit Set) bool {
	return s.Red > limit.Red || s.Green > limit.Green || s.Blue > limit.Blue
}
