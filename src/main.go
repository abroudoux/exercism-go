package main

func TotalBirdCount(birdsPerDay []int) int {
	birdsCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
        birdsCount += birdsPerDay[i]
    }

    return birdsCount
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsCount := 0
	for i := 1; i < week * 7 + 1; i++ {
		birdsCount += birdsPerDay[i - 1]
	}

	return birdsCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	panic("Please implement the FixBirdCountLog() function")
}

func main() {
	list := []int{1, 2, 3}
	println(list[0])
}