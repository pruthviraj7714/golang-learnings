package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func printNumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}

func printAlphabets() {
	for char := 'a'; char <= 'z'; char++ {
		fmt.Println(string(char))
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// go keyword is used to start a goroutine
	// go keyword runs function concurrently
	go sayHello()

	// without go keyword, it will run in the main thread
	// sayHello()

	// without sleep, main function will exit before the goroutine finishes
	time.Sleep(1 * time.Second)

	//multiple goroutines
	//note that the output order is not guaranteed
	//Always pass i as parameter (avoid closure bug) - because of loop variable is shared
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	//anonymous go routines
	go func() {
		fmt.Println("Running async")
	}()

	go printAlphabets()
	go printNumbers()

	time.Sleep(1 * time.Hour)
	fmt.Println("Main function")
}
