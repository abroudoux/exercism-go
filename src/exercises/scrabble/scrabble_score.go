package scrabble

import (
	"strings"
)

func Score(word string) int {
	var score int

	one := "AEIOULNRST"
	two := "DG"
	three := "BCMP"
	four := "FHVWY"
	five := "K"
	eight := "JX"
	ten := "QZ"

	for i := 0; i < len(word); i++ {
		letter := strings.ToUpper(string(word[i]))

		if strings.Contains(one, letter) {
			score = score + 1
		} else if strings.Contains(two, letter) {
			score = score + 2
		} else if strings.Contains(three, letter) {
			score = score + 3
		} else if strings.Contains(four, letter) {
			score = score + 4
		} else if strings.Contains(five, letter) {
			score = score + 5
		} else if strings.Contains(eight, letter) {
			score = score + 8
		} else if strings.Contains(ten, letter) {
			score = score + 10
		}
	}

	return score
}
