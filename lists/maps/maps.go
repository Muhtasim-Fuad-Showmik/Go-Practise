package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
		"Amazon":   "https://amazon.com",
	}
	fmt.Println(websites)

	// Print a specific value from the map
	fmt.Println(websites["Amazon"])

	// Add a new key value to the map
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	// Delete a key from the map
	delete(websites, "Linkedin")
	fmt.Println(websites)
}
