package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var workingCarsPerHour = float64(productionRate) * successRate / 100

	return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCarsPerMinute = CalculateWorkingCarsPerHour(productionRate, successRate) / 60

	return int(workingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupsOfTenCars = int(carsCount / 10)
	var remainingCars = carsCount % 10

	var costPerGroupOfTen = 95000
	var unitCost = 10000

	var cost = uint(groupsOfTenCars*costPerGroupOfTen) + uint(remainingCars*unitCost)

	return cost
}
