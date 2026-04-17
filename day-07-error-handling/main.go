package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// custom errors
func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, fmt.Errorf("insufficient balance: have %d, need %d", balance, amount)
	}
	return balance - amount, nil
}
func doSomething() error {
	return fmt.Errorf("something went wrong")
}

// wrapping errors
func process() error {
	err := doSomething()

	if err != nil {
		return fmt.Errorf("Process failed: %w", err)
	}

	return nil
}

func main() {
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result is: ", result)
	}

	result1, err := withdraw(100, 101)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1)
	}

	result2, err := withdraw(50, 25)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}

	error1 := process()

	if error1 != nil {
		fmt.Println(error1)
	}

	//checking specific error
	error2 := errors.New("not found")

	if error2.Error() == "not found" {
		fmt.Println("Error is not found")
	}

	//errors.Is()

	error3 := errors.New("not found")

	if errors.Is(error2, error3) {
		fmt.Println("Error is not found")
	}

}
