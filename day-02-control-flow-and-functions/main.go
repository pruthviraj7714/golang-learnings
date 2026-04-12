package main

import "fmt"

func main() {

	//if-else
	age := 20

	if age > 18 {
		fmt.Println("You are adult")
	} else {
		fmt.Println("You are not adult")
	}

	//if with initialization
	if score := 20; score > 50 { //score exists only inside if block
		fmt.Println("You won")
	} else {
		fmt.Println("You lost")
	}

	//if-else if-else
	num := 10

	if num > 10 {
		fmt.Println("Number is greater than 10")
	} else if num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is equal to 10")
	}

	//switch
	//No break needed
	//Go automatically breaks (unlike JS)
	day := 2

	switch day {
	case 1:
		fmt.Println("It's Monday")
	case 2:
		fmt.Println("It's Tuesday")
	case 3:
		fmt.Println("It's Wednesday")
	case 4:
		fmt.Println("It's Thursday")
	case 5:
		fmt.Println("It's Friday")
	case 6:
		fmt.Println("It's Saturday")
	case 7:
		fmt.Println("It's Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch without condition
	val := 25

	switch {
	case val > 25:
		fmt.Println("val is greater")
	case val < 25:
		fmt.Println("val is smaller")
	default:
		fmt.Println("val is equal")
	}

	//loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for as while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//infinite loop
	for {
		fmt.Print("Hello")
		break
	}

	//for with continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	//for with range
	for i := range 10 {
		fmt.Println(i)
	}

	//functions
	fmt.Println(add(10, 20))

	//multiple return values
	q, r := divide(15, 3)
	fmt.Println(q, r)

	//named return values
	name, age := getUser()
	fmt.Println(name, age)

	fmt.Println(checkNumberIsEvenOrOdd(14))
}

func add(a, b int) int {
	return a + b
}

// multiple return values
func divide(a, b int) (int, int) {
	return a / b, a % b
}

//named return values

func getUser() (name string, age int) {
	name = "John"
	age = 18
	return //no need to return explicitly
}

func checkNumberIsEvenOrOdd(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}
