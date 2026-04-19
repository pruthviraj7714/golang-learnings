package main

// chan<- int means send only
func sendData(ch chan<- int, data int) {
	ch <- data
}

// <-chan int means receive only
func recieveData(ch <-chan int) {
	val := <-ch
	println(val)
}

func main() {
	ch := make(chan int)
	ch1 := make(chan int)

	go func() {
		ch <- 10 //send value to channel
	}()

	val := <-ch //receive value from channel

	println(val)

	//multiple values
	go func() {
		ch1 <- 11
		ch1 <- 12
	}()
	println(<-ch1)
	println(<-ch1)

	//loop with channels

	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch2 <- i
		}
		close(ch2) // close the channel - no more data will be sent
	}()

	//receive values from channel
	for val := range ch2 {
		println(val)
	}

	//accessing closed channel will cause panic
	// println(<-ch2)

	ch4 := make(chan int)

	go sendData(ch4, 15)
	go sendData(ch4, 17)

	recieveData(ch4)
	recieveData(ch4)
}
