package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    carsPerMinute := float64(productionRate) / 60.0
    workingCarsPerMinute := carsPerMinute * (successRate / 100.0)
    return int(workingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const priceStackCars uint = 95000
    const priceOneCar uint = 10000
    
    var stackCars uint = uint(carsCount) / 10
    var remainingCars uint = uint(carsCount) % 10

    return (priceStackCars * stackCars) + (priceOneCar * remainingCars)
}
