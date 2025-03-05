package darts

func Score(x, y float64) int {
	radius := (x*x + y*y)

	switch {
	case radius <= 1:
		return 10
	case radius <= 25:
		return 5
	case radius <= 100:
		return 1
	default:
		return 0
	}
}
