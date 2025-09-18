package lasagna

func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}
	return len(layers) * prepTimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodleCount := 0
	sauceCount := 0
	for _, v := range layers {
		if v == "sauce" {
			sauceCount++
		} else if v == "noodles" {
			noodleCount++
		}
	}
	return noodleCount * 50, float64(sauceCount) * 0.2
}

func AddSecretIngredient(friendIngredients, ownIngredients []string) {
	ownIngredients[len(ownIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

func ScaleRecipe(amountsForTwo []float64, portions int) []float64 {
	amountsNeeded := make([]float64, len(amountsForTwo))
	multiplier := portions / 2
	mod := portions % 2
	for i, _ := range amountsNeeded {
		amountsNeeded[i] = amountsForTwo[i]*float64(multiplier) + float64(mod)*amountsForTwo[i]/2
	}
	return amountsNeeded
}
