package exercises

// I have to comment this line because this function is already defined in gophers_gorgeous_lasagna
// func PreparationTime(layers []string, avgTimePerLayer int) int {
// 	if avgTimePerLayer == 0 {
// 		return len(layers) * 2
// 	}

// 	return len(layers) * avgTimePerLayer
// }

func Quantities(layers []string) (int, float64) {
	countNoodles := countOccurences(layers, "noodles")
	countSauce := countOccurences(layers, "sauce")

	return countNoodles * 50, float64(countSauce) * 0.2
}

func countOccurences(slice []string, ref string) int {
	occ := 0
	for _, v := range slice {
		if v == ref {
			occ++
		}
	}

	return occ
}

func AddSecretIngredient(friendList []string, ingredientList []string) []string {
	lastIngredientOfFriendList := friendList[len(friendList) - 1]
	ingredientListSliced := ingredientList[:len(ingredientList) - 1]
	return append(ingredientListSliced, lastIngredientOfFriendList)
}

func ScaleRecipe(amounts []float64, quantities int) []float64 {
	var scaledAmounts []float64
	for _, recipeQty := range amounts {
		scaledAmounts = append(scaledAmounts, recipeQty / 2.00 * float64(quantities))
	}
	return scaledAmounts
}
