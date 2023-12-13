package day6

func CalculatePart1(races []Race) int {
	result := 1
	for _, race := range races {
		result *= race.numberOfWaysToWin()
	}
	return result
}

type Race struct {
	TimeMs   int
	Distance int
}

const VELOCITY = 1

func (r Race) numberOfWaysToWin() int {
	result := 0
	for holdDownButtonMs := 0; holdDownButtonMs < r.TimeMs; holdDownButtonMs++ {
		totalDistance := (r.TimeMs - holdDownButtonMs) * holdDownButtonMs * VELOCITY
		if totalDistance > r.Distance {
			result += 1
		}
	}
	return result
}
