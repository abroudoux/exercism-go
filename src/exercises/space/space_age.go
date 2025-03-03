package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	earthSeconds := 31557600.0

	switch planet {
	case "Mercury":
		return seconds / (0.2408467 * earthSeconds)
	case "Venus":
		return seconds / (0.61519726 * earthSeconds)
	case "Earth":
		return seconds / 31557600
	case "Mars":
		return seconds / (1.8808158 * earthSeconds)
	case "Jupiter":
		return seconds / (11.862615 * earthSeconds)
	case "Saturn":
		return seconds / (29.447498 * earthSeconds)
	case "Uranus":
		return seconds / (84.016846 * earthSeconds)
	case "Neptune":
		return seconds / (164.79132 * earthSeconds)
	}

	return -1.0
}
