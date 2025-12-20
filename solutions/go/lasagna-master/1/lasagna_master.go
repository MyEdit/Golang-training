package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTimeInMin int) int {
	if avgTimeInMin <= 0 {
        return len(layers) * 2
    }
    
    return len(layers) * avgTimeInMin
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var sauces float64 = 0.0
    var noodles int = 0
    
    for i := 0; i < len(layers); i++ {
        switch layers[i] {
            case "sauce":
            	sauces += 0.2
            case "noodles":
            	noodles += 50
        }
    }

    return noodles, sauces
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    newQuantities := make([]float64, len(quantities))
    
    for i := 0; i < len(quantities); i++ {
    	newQuantities[i] = quantities[i] * (float64(portions) / 2)
    }

    return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
