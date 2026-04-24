package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status string
	Err    error
}

func worker(ctx context.Context, id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d is fetching %s\n", id, url)

		//Create a new request with the context.
		req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

		//Create a new client.
		client := http.Client{}

		//Do() function is used to send the request to the server.
		resp, err := client.Do(req)

		//Check if there is an error.
		if err != nil {
			results <- Result{URL: url, Err: err}
			continue
		}

		//Send the result to the results channel.
		results <- Result{
			URL:    url,
			Status: resp.Status,
		}

		//Close the response body.
		resp.Body.Close()

	}

}

func main() {

	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://invalid-url",
	}

	//Make a channel for jobs and a channel for results.
	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))

	//Context is used to cancel the request if it takes too long.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	//Cancel() function is called when the function returns.
	defer cancel()

	var wg sync.WaitGroup

	numWorkers := 3

	//Start the workers.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}

	//Send jobs to the jobs channel.
	for _, url := range urls {
		jobs <- url
	}

	//Close the jobs channel.
	close(jobs)

	//Wait for all the workers to finish.
	wg.Wait()

	//Close the results channel.
	close(results)

	//Print the results.
	for res := range results {
		if res.Err != nil {
			fmt.Printf("Error fetching %s: %v\n", res.URL, res.Err)
		} else {
			fmt.Printf("Fetched %s: %s\n", res.URL, res.Status)
		}
	}

}
