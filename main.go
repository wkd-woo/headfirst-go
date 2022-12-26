package main

import "fmt"

func main() {
	price := 100
	fmt.Println("Price is ", price, "dollors")

	taxRate := 0.08
	tax := price * int(taxRate)

	fmt.Println("Tax is ", tax, "dollors.")
	total := price + tax
	fmt.Println("Total cost is ", total)

	availableFunds := 120
	fmt.Println(availableFunds, "dollors available")
	fmt.Println("Within budget?", total <= availableFunds)
}
