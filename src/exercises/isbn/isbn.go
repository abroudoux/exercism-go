package isbn

import "regexp"

func IsValidISBN(isbn string) bool {
	if len(isbn) != 13 {
		return false
	}

	pattern := `^\d-\d{3}-\d{4}-[\dX]$`
	if !regexp.MustCompile(pattern).MatchString(isbn) {
		return false
	}

	return true
}
