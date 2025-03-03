package luhn

import (
	"regexp"
	"strings"
)

func Valid(id string) bool {
	idTrimed := strings.ReplaceAll(id, " ", "")
	if len(idTrimed) <= 1 {
		return false
	}

	var re = regexp.MustCompile(`^[0-9]+$`)

	if !re.MatchString(idTrimed) {
		return false
	}

	if len(idTrimed)%2 == 0 {

	} else {
	}

	return true
}
