package main

import (
	"fmt"
	"time"
)

func worker(name string, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- name + " done"
}

func main() {

	//unbuffered channel
	// ch := make(chan int)
	// ch <- 10 // ❌ blocks until someone receives

	//buffered channel
	ch := make(chan int, 2) //capacity 2

	ch <- 10
	ch <- 12

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3 // ❌ blocks (buffer full)

	fmt.Println(<-ch)

	//select
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "from ch1"
	}()

	go func() {
		ch2 <- "from ch2"
	}()

	//👉 select picks whichever channel is ready first
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	//if no channel is ready, default case is executed
	default:
		fmt.Println("No Blocking")
	}

	//timeout pattern

	ch3 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "done"
	}()

	select {
	case msg := <-ch3:
		fmt.Println(msg)
	//time.After returns a channel that receives a value after the specified duration
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	//multiple channel handling
	ch4 := make(chan string)
	ch5 := make(chan string)

	go func() {
		ch4 <- "message from ch4"
	}()

	go func() {
		ch5 <- "message from ch5"
	}()

	//infinite loop to receive values from channels
	// for {
	// 	select {
	// 	case msg1 := <-ch4:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-ch5:
	// 		fmt.Println(msg2)
	// 	}
	// }

	ch6 := make(chan string)
	ch7 := make(chan string)

	go worker("worker1", ch6)
	go worker("worker2", ch7)

	//receive values from channels fixed iterations
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch6:
			fmt.Println(msg1)
		case msg2 := <-ch7:
			fmt.Println(msg2)
		}
	}

	fmt.Println("All work is done")
}
