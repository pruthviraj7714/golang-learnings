package main

import "fmt"

type Address struct {
	City string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func (u User) greet() {
	fmt.Println("\nHello", u.Name)
}

func (u *User) incrementAge() { //pointer receiver
	u.Age++
}

func (u User) fullInfo() {
	fmt.Printf("The User with name %s and age %d lives in %s", u.Name, u.Age, u.Address.City)
}

func main() {
	u1 := User{
		Name: "Tony Stark",
		Age:  20,
		Address: Address{
			City: "Noida",
		},
	}

	fmt.Printf("The Age of %s is %d", u1.Name, u1.Age)

	u1.greet()

	u1.incrementAge()
	u1.incrementAge()
	u1.incrementAge()
	u1.incrementAge()

	fmt.Printf("The Age of %s is %d\n", u1.Name, u1.Age)

	//struct composition
	u1.fullInfo()

	//anonymous struct
	a := struct {
		Name string
		Age  int
	}{
		Name: "Tony Stark",
		Age:  20,
	}
	fmt.Println(a)

}
