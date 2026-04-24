# Day 14 - Concurrent URL Fetcher

## Description

A concurrent URL fetcher built in Go using worker pool pattern.

## Features

- Concurrent HTTP requests
- Worker pool (controlled concurrency)
- Context-based timeout
- Error handling

## Tech Used

- Goroutines
- Channels
- WaitGroup
- Context

## 📝 Notes

- Built real concurrency system
- Understood worker pool pattern deeply
- Applied timeout handling
- used channels for communication between goroutines
- WaitGroup is used to wait for all the goroutines to finish.
- Context is used to cancel the request if it takes too long.
