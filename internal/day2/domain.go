package day2

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
