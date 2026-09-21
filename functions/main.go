package main

import "fmt"

// For complex functions defining custom function types can save characters
// from being repeatedly typed and maintained
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
	
	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)
	
	transformedNumbers ::= transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers ::= transformNumbers(&moreNumbers, transformerFn2)
	
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

// Get function as a return value
func getTransformerFunction(numbers *[]int) transformFn {
	if {*numbers}[0] == 1 {
		return double
	} else {
		return triple
	}
}

// Passing functions as values into the parameters for another function
func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}

	return tNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
