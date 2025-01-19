package main

func TotalBirdCount(birdsPerDay []int) int {
	birdsCount := 0

	for i := 0; i < len(birdsPerDay); i++ {
        birdsCount += birdsPerDay[i]
    }

    return birdsCount
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsToCount := birdsPerDay[week - 1: week - 1 +7]
	birdsCount := 0

	for i := 0; i < len(birdsToCount); i++ {
		birdsCount += birdsToCount[i]
	}

	return birdsCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	panic("Please implement the FixBirdCountLog() function")
}

func main() {
	weekCount := 1
	allWeeks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	week := allWeeks[]

	for i := 0; i < len(first); i++ {
		println(first[i])
	}
}