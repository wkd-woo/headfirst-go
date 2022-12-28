package main

import "fmt"

var metersPerLiter float64

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / metersPerLiter
}

func main() {
	metersPerLiter = 10.0

	total := 0.0
	amount := paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	fmt.Printf("Total: %0.2f liters\n", total)
}
