package main

import "fmt"

func main() {
	// 1. Initialization:
	colors := map[string]string{
		"red": "#red",
		"green": "#green",
		"black": "#black",
	}

	// 2. Initialization: zero value nil
	// var colors map[string]string

	// 3. Initialization: make function
	// colors := make(map[string]string)

	// 4. Insert value
	// colors["white"] = "#white"

	// 5. Retrieve
	// var x = colors["white"]
	// fmt.Println(x)

	// 6. Delete
	// delete(colors,"white")

	// 7. check is key is present
	// v, ok := colors["white"]
	// fmt.Println("The value:", v, "Present?", ok)
	// fmt.Println(colors)

	// 8. Iterate
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
