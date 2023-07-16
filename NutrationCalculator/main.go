package main

import (
	"fmt"
)

func main() {

	ns := GetNutritionalScore(
		NutritionalData{
			Energy:             EnergyFromKcal(34),
			Sugars:             SugarGram(100),
			SatuatedFattyAcids: SaturatedFattyAcids(2),
			Sodium:             SodiumMilligram(6),
			Fruits:             FruitsPercent(3),
			Fiber:              FibreGram(9),
			Protein:            ProteinGram(3),
		}, Food)

	fmt.Printf("Nutritional Score :%d \n", ns.Value)
	// fmt.Printf("NutriScore Score :%d \n", ns.GetNutriScore())

}
