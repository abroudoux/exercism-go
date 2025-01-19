package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	birdsCount := 0

	for i := 0; i < len(birdsPerDay); i++ {
        birdsCount += birdsPerDay[i]
    }

    return birdsCount
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	return TotalBirdCount(birdsPerDay[(week - 1) * 7:week * 7])
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i % 2 == 0 {
			birdsPerDay[i] = v + 1
		}
	}

	return birdsPerDay
}