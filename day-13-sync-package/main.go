package main

import (
	"fmt"
	"sync"
)

var counter int
var data []int
var mu sync.Mutex
var rmu sync.RWMutex

func increment(wg *sync.WaitGroup) {
	//Done() function is called when the goroutine finishes.
	defer wg.Done()
	for i := 0; i <= 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func write(wg *sync.WaitGroup, val int) {
	defer wg.Done()

	//Lock() function is used to lock the RWMutex for writing.
	rmu.Lock()
	fmt.Println("Writing Counter:", val)
	data = append(data, val)
	//Unlock() function is used to unlock the RWMutex for writing.
	rmu.Unlock()
}

func read(wg *sync.WaitGroup) {
	defer wg.Done()

	//RLock() function is used to lock the RWMutex for reading.
	rmu.RLock()
	fmt.Println("Reading Counter:", data)
	//RUnlock() function is used to unlock the RWMutex for reading.
	rmu.RUnlock()
}

func initConfig() {
	fmt.Println("Configuration initializing")
}

func main() {
	//WaitGroup is used to wait for all the goroutines to finish.
	var wg sync.WaitGroup
	var once sync.Once

	//Do() function is used to execute the function only once.
	//We are creating 5 goroutines and each goroutine is calling the initConfig() function.
	//But the initConfig() function will be executed only once because of the once.Do() function.
	for i := 0; i < 5; i++ {
		go func() {
			once.Do(initConfig)
		}()
	}

	//Add() function is used to add the number of goroutines to wait for.
	wg.Add(7)

	//goroutine are started here and wg.Done() is called when the goroutine finishes.
	go increment(&wg)
	go increment(&wg)
	go increment(&wg)

	go write(&wg, 101)
	go read(&wg)
	go write(&wg, 102)
	go read(&wg)

	//Wait() function is used to wait for all the goroutines to finish.
	wg.Wait()

	fmt.Println("Counter:", counter)
}
