package main

import (
	"../advent"
	"fmt"
	"math"
)

func fuelNeeded(moduleMass int) int {
	return int(math.Floor(float64(moduleMass/3))) - 2
}

// TotalFuelNeeded returns the total fuel needed
func TotalFuelNeeded(modules []int) int {
	totalFuel := 0
	for _, mass := range modules {
		moduleFuel := fuelNeeded(mass)
		additionalFuel := fuelNeeded(moduleFuel)
		for {
			if additionalFuel <= 0 {
				break
			}
			moduleFuel += additionalFuel
			additionalFuel = fuelNeeded(additionalFuel)
		}
		totalFuel += moduleFuel
	}
	return totalFuel
}

func main() {
	input := advent.ReadIntArrayInput()
	fmt.Println(TotalFuelNeeded(input))
}
