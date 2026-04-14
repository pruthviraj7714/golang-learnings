package main

import "fmt"

// without pointer
func increment(x int) {
	x++
}

// with pointer
func incrementPointer(x *int) {
	*x++
}

type User struct {
	Name string
	Age  int
}

func (u *User) incrementAge() {
	u.Age++
}

func main() {

	x := 10
	var p *int = &x

	fmt.Println("The value of x is:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Pointer p:", p)
	fmt.Println("Value via pointer:", *p)

	increment(x)
	fmt.Println("After increment(x):", x) // x is still 10 because it was passed by value

	incrementPointer(&x)
	fmt.Println("After incrementPointer(&x):", x) // x is 11 because it was passed by reference

	user1 := User{
		Name: "Tony",
		Age:  25,
	}

	user1.incrementAge() //Go automatically handles &user for you here

	fmt.Println("After incrementAge:", user1.Age)

	var p2 *int
	fmt.Println(p2) //nil

	// fmt.Println(*p2) //panic: runtime error: invalid memory address or nil pointer dereference
}
