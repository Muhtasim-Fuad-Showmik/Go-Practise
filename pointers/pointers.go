package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age: ", age)
	fmt.Println("Age Pointer: ", agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult Years: ", adultYears)

	updateAge(agePointer)
	fmt.Println("Age after update: ", age)
}

func getAdultYears(age *int) int {
	return *age - 18
}

func updateAge(age *int) {
	*age = 33
}
