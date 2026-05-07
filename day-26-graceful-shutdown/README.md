# Day 26 - Graceful Shutdown

## What I Built

- Graceful shutdown for Go server
- Signal handling (SIGINT, SIGTERM)
- Timeout-based shutdown

## Key Learnings

- http.Server shutdown
- Context for lifecycle management
- Clean resource handling

## What's Happening

Step-by-step:

- Server runs in goroutine
- Wait for signal
- Signal received → shutdown starts
- Context gives 5s timeout
- Server stops gracefully

## Improvements

- Safe shutdown
- No abrupt termination

## Quick Notes

### 👉 What is graceful shutdown?

Cleanly stopping server without killing in-progress work

### 👉 Why needed?

prevent data loss
improve reliability

### 👉 How in Go?

http.Server.Shutdown
context.WithTimeout
signal handling
