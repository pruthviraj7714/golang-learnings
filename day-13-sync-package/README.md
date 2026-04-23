# Day 13 - Sync Package (Mutex, RWMutex)

## What I Learned

- Race conditions
- Mutex for safe concurrent access
- RWMutex for read-heavy workloads
- WaitGroup for synchronization

## 📝 Notes

- Mutex ensures safe access to shared memory in concurrent programs.
- WaitGroup is used to wait for all the goroutines to finish.
- RWMutex is used for read-heavy workloads.
  - RWMutex is more performant than Mutex for read-heavy workloads.
  - RWMutex is less performant than Mutex for write-heavy workloads.
- Atomic operations are used for safe concurrent access to shared memory.
- Once is used to execute a function only once.

## Practice

- Thread-safe counter implementation
