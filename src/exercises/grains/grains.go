package grains

import "errors"

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("please choose a valid number")
	}

	return uint64(1 << (number - 1)), nil
}

func Total() uint64 {
	var sum int
	for i := 0; i <= 64; i++ {
		grains, _ := Square(i)
		sum += int(grains)
	}

	return uint64(sum)
}
