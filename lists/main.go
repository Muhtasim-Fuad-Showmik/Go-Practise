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
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	// Slicing an array in the end
	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)

	// Update all references to array data by updating slices
	featuredPrices[0] = 199.99
	fmt.Println(prices)

	// Array operations
	fmt.Println(len(featuredPrices), cap(featuredPrices))

	highlightedPrices = highlightedPrices[1:3]
	fmt.Println("Highlighted Prices:")
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	fmt.Println("Dynamic Arrays:")
	// Array without any predefined length restriction
	dynamicPrices := []float64{10.99, 6.99}

	// Adding new values to the array
	updatedPrices := append(dynamicPrices, 5.99)

	fmt.Println(updatedPrices)
}
