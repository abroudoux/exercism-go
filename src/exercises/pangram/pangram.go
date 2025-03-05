package pangram

import (
	"slices"
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	if len(input) < 26 {
		return false
	}

	inputTrimed := strings.ReplaceAll(input, " ", "")
	if len(inputTrimed) < 26 {
		return false
	}

	var allLetters []rune
	for _, c := range strings.ToUpper(inputTrimed) {
		if !unicode.IsLetter(c) {
			continue
		}

		if !slices.Contains(allLetters, c) {
			allLetters = append(allLetters, c)
		}
	}

	return len(allLetters) == 26
}
