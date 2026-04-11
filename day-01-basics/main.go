package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// variables

	var name string = "Tony"
	age := 22 // short variable declaration

	const pi = 3.14

	//Data types
	var isStudent bool = true
	var height float64 = 5.9
	var grade rune = 'A' // rune is an alias for int32 and represents a Unicode code point
	var count int = 10

	fmt.Printf("Name: %s, Age: %d\n", name, age) // %s for string, %d for int
	fmt.Printf("Is Student: %t\n", isStudent)    // %t for boolean
	fmt.Printf("Height: %f\n", height)           // %f for float
	fmt.Printf("Grade: %c\n", grade)             // %c for rune
	fmt.Printf("Count: %d\n", count)             // %d for int

	// constants
	fmt.Println("Pi: ", pi)

	fmt.Printf("Type of age: %T\n", age) // %T for type
}
