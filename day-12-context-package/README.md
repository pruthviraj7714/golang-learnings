# Day 12 - Context Package

## What I Learned

- Context for cancellation and timeouts
- WithTimeout and WithCancel
- Passing context in functions
- Graceful shutdown patterns

## Notes 🗒️

- Context is essential for building production-grade Go services.
- It is used to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between goroutines.
- Root context: context.Background()
- Context with timeout: context.WithTimeout(parent, duration)
- Context with cancel: context.WithCancel(parent)
- Context values: context.WithValue(parent, key, value)
- Context is immutable
- It is passed as the first argument to functions that need it
- It is used to propagate cancellation signals and request-scoped values
- Use context.WithCancel to cancel the context manually
- Use context.WithTimeout to cancel the context after a specific duration
- Use context.WithValue to pass request-scoped values to the child context
- Always defer the cancel function to prevent context leaks
- Do not pass values in context
- Do not use context for passing data between goroutines

## Practice

- Worker with timeout cancellation
