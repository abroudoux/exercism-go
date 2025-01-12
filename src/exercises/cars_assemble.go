package exercises

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

func CalculateCost(carsCount int) uint {
	groupsCount := carsCount / 10
	solosCars := carsCount % 10
	groupsCost := 95.000 * uint(groupsCount)
	solosCarsCost := 10.000 * uint(solosCars)
	return (groupsCost + solosCarsCost) * 1000
}
