package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// Array Definitions
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	var productNames [4]string = [4]string{"A Book"}

	// Array value assignment at an index
	productNames[2] = "A Carpet"

	// Output entire arrays
	fmt.Println(prices)
	fmt.Println(productNames)

	// Output first value of the array
	fmt.Println(prices[0])

	// Slicing an array in the middle
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	// Slicing an array in the start
	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	// Slicing an array in the end
	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)
}
