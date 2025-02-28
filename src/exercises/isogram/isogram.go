package isogram

import (
	"slices"
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	var letters []string

	for i := 0; i < len(word); i++ {
		char := strings.ToUpper(string(word[i]))

		if slices.Contains(letters, char) {
			return false
		}

		if unicode.IsLetter(rune(word[i])) {
			letters = append(letters, char)
		}
	}

	return true
}
