# Day 10 - Buffered Channels & Select

## What I Learned

- Buffered vs unbuffered channels
- Select statement for handling multiple channels
- Non-blocking operations using default
- Timeout patterns using time.After

## 📝 Notes

- Unbuffered channels: these channels are created using `make(chan int)` and they block until the sender and receiver are ready.

- Buffered channels: these channels are created using `make(chan int, capacity)` and they block only when the buffer is full.
  - Allow storing values without immediate receiver

- Select allows handling multiple concurrent operations efficiently.

- Default case in select is executed when no channel is ready.

- Time after returns a channel that receives a value after the specified duration.

- Timeout Pattern:
  👉 Real-world use:
  - API timeouts
  - DB query timeout
  - network calls

- Multiple Channel Handling
  - Infinite loop to receive values from channels
    👉 Used in:
    - worker systems
    - event processing
    - message brokers
- Go vs Node Insight
- Node:

```javascript
Promise.race([task1, task2]);
```

- Go:

```go
select {
case <-ch1:
case <-ch2:
}
```

👉 select = Promise.race + event loop control

## Practice

- Worker example with select
- Timeout handling
