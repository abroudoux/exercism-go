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

	sum := 0
	isSecond := false

	for i := len(idTrimed) - 1; i >= 0; i-- {
		d := int(idTrimed[i] - '0')

		if isSecond {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		sum += d
		isSecond = !isSecond
	}

	return sum%10 == 0
}
