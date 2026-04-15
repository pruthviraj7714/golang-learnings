package main

import "fmt"

func main() {

	//Arrays - Rarely used in real apps
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr[0])

	//slices
	nums := []int{1, 2, 3}

	nums = append(nums, 5) // append is used to add elements to a slice

	fmt.Println(nums[3])

	sub := nums[1:3] // index 1 to 2

	fmt.Println(sub)

	fmt.Println(len(nums)) // length
	fmt.Println(cap(nums)) // capacity

	//Slices are reference types
	a := []int{1, 2, 3}
	b := a

	b[0] = 100

	fmt.Println(a) //[100 2 3] Reason: Slices are reference types

	//Fix
	c := make([]int, len(a)) // Creates a new slice with the same length as a
	copy(c, a)               // Copies the elements of a to c
	fmt.Println(c)

	//Maps

	m := make(map[string]int)

	m["apple"] = 10
	m["banana"] = 20
	m["sdfa"] = 20

	fmt.Println(m["apple"])

	//check if key exists

	val, ok := m["apple"]

	if ok {
		fmt.Println("Exists:", val)
	} else {
		fmt.Println("Not found")
	}

	//delete
	delete(m, "apple")
	fmt.Println(m)

	//Iterate over map
	for key, value := range m {
		fmt.Println(key, value)
	}

	//loop over slice
	for i, v := range nums {
		fmt.Println(i, v)
	}

	//loop over array
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
