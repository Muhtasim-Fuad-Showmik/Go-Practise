package main

import "fmt"

type product struct {
	id    string
	title string
	price float64
}

func main() {
	// Create array of 3 hobbies and output them
	hobbies := [3]string{"Gaming", "Coding", "Reading"}
	fmt.Println(hobbies)

	// Output first element standalone
	fmt.Println(hobbies[0])

	// Output 2nd and 3rd element as a separate slice
	fmt.Println(hobbies[1:3])

	// Create a brand new slice with the first and second element
	// mainHobbies := hobbies[0:2]
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies, cap(mainHobbies))

	// Output 2nd and 3rd element using your earlier slice
	seondaryHobbies := mainHobbies[1:3]
	fmt.Println(seondaryHobbies, cap(seondaryHobbies))

	// Create a dynamic array of goals
	courseGoals := []string{"Learn Go", "Master Go"}
	fmt.Println(courseGoals)

	// Replace second goal
	courseGoals[1] = "Learn all the details"

	// Add a 3rd goal
	courseGoals = append(courseGoals, "Learn all the essentials")

	fmt.Println(courseGoals)

	// Create a dynamic list of products
	products := []product{
		{
			id:    "P1",
			title: "Computer",
			price: 1000.0,
		},
		{
			id:    "P2",
			title: "Mouse",
			price: 5.54,
		},
	}
	fmt.Println(products)

	// Add a third product to the dynamic array
	newProduct := product{
		id:    "P3",
		title: "Keyboard",
		price: 10.99,
	}
	products = append(products, newProduct)

	fmt.Println(products)
}
