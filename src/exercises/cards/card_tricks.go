package cards

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(slice []int, index int) int {
	if index < 0 || index > len(slice) {
		return -1
	}

	return slice[index]
}

func SetItem(slice []int, index, value int) []int {
	if index > len(slice) - 1 || index < 0 {
		return append(slice, value)
	}

	slice[index] = value
	return slice
}

func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	}

	return append(values, slice...)
}

func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice) {
		return slice
	}

	return append(slice[:index], slice[index + 1:]...)
}
