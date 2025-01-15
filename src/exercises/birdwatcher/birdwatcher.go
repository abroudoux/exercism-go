package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	birdsCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
        birdsCount += birdsPerDay[i]
    }

    return birdsCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	return TotalBirdCount(birdsPerDay[week * 7 - 1:week * 7 + 7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	panic("Please implement the FixBirdCountLog() function")
}
