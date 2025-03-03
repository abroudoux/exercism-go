package resistorcolor

func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

func ColorCode(color string) int {
	colors := Colors()

	for i := 0; i < len(colors); i++ {
		if colors[i] == color {
			return i
		}
	}

	return -1
}
