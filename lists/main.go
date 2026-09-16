package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// Create an array with length 2 (with both elements being set to null)
	// and allocating 5 slots for the array
	userNames := make([]string, 2, 5)

	// Replace null with actual values
	userNames[0] = "Julie"
	userNames[1] = "John"

	// Add more values to the array
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	// Make a map with 3 spaces allocated for it
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 5.7
	courseRatings[".NET"] = 7.7

	courseRatings.output()
}
