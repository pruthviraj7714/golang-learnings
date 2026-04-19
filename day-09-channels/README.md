# Day 09 - Channels in Go

## What I Learned

- Channels for communication between goroutines
- A channel is a pipe to communicate between goroutines.
  - Send data ➡️
  - Receive data ⬅️
  - Synchronization built-in 🔥
- Channels are blocking by default
  - Send blocks until someone receives
  - Receive blocks until someone sends
- Sending and receiving data
- Closing channels
- Iterating over channels
- Think:
  - Goroutine = worker
  - Channel = queue
  - Similar to:
    - Kafka topic
    - Redis queue
- Node vs Go
  - Node: Async with callbacks/promises
  - Go: Goroutines + channels
  - Go is closer to: “Communicating Sequential Processes (CSP)”

## 📝 Notes

- Channels enable safe communication and synchronization between goroutines.
- Channels are typed.
- Code snippets

```go
// create a channel
ch := make(chan int)

// send value to channel
ch <- 10

// receive value from channel
val := <-ch

// close the channel
close(ch)

// iterate over channel
for val := range ch {
	println(val)
}
```
