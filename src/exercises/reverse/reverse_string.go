package reverse

func Reverse(input string) string {
	var stringReversed []rune

	for _, r := range input {
		stringReversed = append([]rune{r}, stringReversed...)
	}

	return string(stringReversed)
}
