package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meow!"
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

type PaymentProcessor interface {
	Pay(amount int) string
}

type Stripe struct{}

func (s Stripe) Pay(amount int) string {
	return fmt.Sprintf("Paid %d using Stripe", amount)
}

type Razorpay struct{}

func (r Razorpay) Pay(amount int) string {
	return fmt.Sprintf("Paid %d using razorpay", amount)
}

func processPayment(p PaymentProcessor, amount int) {
	fmt.Println(p.Pay(amount))
}

func checkType(x interface{}) {
	switch y := x.(type) {
	case int:
		fmt.Println("it is int", y)
	case string:
		fmt.Println("it is string", y)
	case bool:
		fmt.Println("it is bool", y)
	case float64:
		fmt.Println("it is float64", y)
	default:
		fmt.Println("it is unknown type")
	}
}

func main() {
	dog := Dog{}
	cat := Cat{}

	makeSound(dog)
	makeSound(cat)

	stripe := Stripe{}
	razorpay := Razorpay{}

	processPayment(stripe, 100)
	processPayment(razorpay, 200)

	//empty interfaces - can hold any type (like any in ts)
	var x interface{}
	x = 10
	fmt.Println(x)
	x = "hello"
	fmt.Println(x)
	x = true
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)

	// Type assertion

	var y interface{} = "hello"
	str := y.(string) // type assertion
	fmt.Println(str)

	str, ok := y.(string)
	if ok {
		fmt.Println("it is string")
	}

	_, ok = y.(int)
	if !ok {
		fmt.Println("it is not int")
	}

	checkType("10")
	checkType(10)
	checkType(true)
	checkType(3.14)
	checkType(Dog{})
}
