package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(1 * time.Second)
		results <- job * 2
	}
}

func main() {
	//problem we are solving

	//1. too many goroutines
	//2. too many resources
	//3. too many database connections
	//4. too many file operations
	//5. too many network requests

	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		fmt.Println("Processing job")
	// 	}()
	// }

	//solution : worker pool

	//Fixed number of workers
	//Jobs go into queue
	//Workers pick jobs

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	//send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= 5; r++ {
		fmt.Println("Result:", <-results)
	}

}
