package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating a map (jjust declaration)
	m := make(map[string]string)

	// adding key-value pairs
	m["name"] = "Alice"
	m["city"] = "Wonderland"
	m["occupation"] = "Adventurer"
	m["hobby"] = "Exploring"

	// alternative way to create and initialize a map
	// m := map[string]string{
	// 	"name":       "Alice",
	// 	"city":       "Wonderland",
	// 	"occupation": "Adventurer",
	// 	"hobby":      "Exploring",
	// }

	// retrieving a value
	name := m["name"]
	city := m["city"]
	occupation := m["occupation"]

	fmt.Println(name, city, occupation)

	// changing a value
	m["city"] = "Looking Glass Land"
	fmt.Println("Updated city:", m["city"])

	// deleting a key-value pair
	delete(m, "occupation")
	fmt.Println("Occupation exists after deletion:", m["occupation"]) // if key-value pair doesn't exist, returns zero value

	// checking if a key exists
	_, exists := m["occupation"]
	fmt.Println("Does occupation exist?", exists)

	fmt.Println(len(m))

	// comparing two maps using maps package
	m1 := map[string]string{
		"name":       "Alice",
		"city":       "Wonderland",
		"occupation": "Adventurer",
		"hobby":      "Exploring",
	}

	fmt.Println(maps.Equal(m, m1)) // output : false ->  because we deleted "occupation" key from m
}
