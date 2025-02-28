package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	var str string = ""

	if number%3 == 0 {
		str = str + "Pling"
	}

	if number%5 == 0 {
		str = str + "Plang"
	}

	if number%7 == 0 {
		str = str + "Plong"
	}

	if str == "" {
		return strconv.Itoa(number)
	}

	return str
}
