# Day 08 - Goroutines

## What I Learned

- What goroutines are
- How to run functions concurrently using `go`
- Basic concurrency in Go
- Why Goroutines are Powerful compare to Node.js
  - Node: Single-threaded (event loop)
  - Go: True concurrency
  - Go: Parallel execution (multi-core)
  - Example use cases:
    - API calls in parallel
    - Worker systems
    - Real-time systems
    - Background jobs
- Right now:
  - We’re using time.Sleep ❌ (not ideal)
- Tomorrow:
  - We fix this using Channels (REAL concurrency control)

## 📝 Notes

- Goroutines are lightweight threads that enable powerful concurrency in Go.
- go keyword is used to start a goroutine
- go keyword runs function concurrently
- without sleep, main function will exit before the goroutine finishes
- multiple goroutines
- note that the output order is not guaranteed
- Always pass i as parameter (avoid closure bug) - because of loop variable is shared
- anonymous go routines

eg. Code Snippets

- Declaring goroutine

```go
go sayHello()
```

- Multiple goroutines

```go
for i := 0; i < 5; i++ {
	go func(i int) {
		fmt.Println(i)
	}(i)
}
```

- Anonymous goroutine

```go
go func() {
	fmt.Println("Running async")
}()
```

## Practice

- Running multiple goroutines
- Printing numbers and alphabets concurrently
