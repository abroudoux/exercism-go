package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("please enter a valid number")
	}

	var roman strings.Builder

	for {
		if input == 0 {
			break
		}

		if input >= 1000 {
			roman.WriteString("M")
			input -= 1000
			continue
		}

		if input >= 900 {
			roman.WriteString("CM")
			input -= 900
			continue
		}

		if input >= 500 {
			roman.WriteString("D")
			input -= 500
			continue
		}

		if input >= 400 {
			roman.WriteString("CD")
			input -= 400
			continue
		}

		if input >= 100 {
			roman.WriteString("C")
			input -= 100
			continue
		}

		if input >= 90 {
			roman.WriteString("XC")
			input -= 90
			continue
		}

		if input >= 50 {
			roman.WriteString("L")
			input -= 50
			continue
		}

		if input >= 40 {
			roman.WriteString("XL")
			input -= 40
			continue
		}

		if input >= 10 {
			roman.WriteString("X")
			input -= 10
			continue
		}

		if input == 9 {
			roman.WriteString("IX")
			input -= 9
			continue
		}

		if input >= 5 {
			roman.WriteString("V")
			input -= 5
			continue
		}

		if input == 4 {
			roman.WriteString("IV")
			input -= 4
			continue
		}

		if input >= 1 {
			roman.WriteString("I")
			input -= 1
			continue
		}
	}

	return roman.String(), nil
}
