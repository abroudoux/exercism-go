package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	var h = make(Histogram)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0

	if d == "" {
		return h, nil
	}

	for _, r := range d {
		if _, ok := h[r]; !ok {
			return Histogram{}, errors.New("invalid character")
		}

		h[r] = h[r] + 1
	}

	return h, nil
}
