package main

import (
	"context"
	"fmt"
	"time"
)

func longTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Long task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

// general pattern
func worker(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("work done")
		return nil
	case <-ctx.Done():
		fmt.Println("cancelled")
		return ctx.Err()
	}
}

func main() {
	// context is used to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

	//root context
	// ctx0 := context.Background()

	//with timeout after 2 sec it will automatically cancel
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)

	//defer the cancel - important to prevent context leak
	defer cancel1()

	longTask(ctx1)

	//manual cancellation
	ctx2, cancel2 := context.WithCancel(context.Background())

	//defer the cancel - important to prevent context leak
	// defer cancel2()

	go longTask(ctx2)
	time.Sleep(1 * time.Second) //wait for 1 second
	cancel2()                   //manually cancel

	//context values

	// Pass request-scoped values to the child (eg user id)
	ctx3 := context.WithValue(context.Background(), "userId", 1233)

	//Retrive value from context
	fmt.Println(ctx3.Value("userId"))

}
