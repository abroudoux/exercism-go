package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	valuesUpdated := make(map[string]int)

	for points, arr := range in {
		for _, letter := range arr {
			valuesUpdated[strings.ToLower(letter)] = points
		}
	}

	return valuesUpdated
}
