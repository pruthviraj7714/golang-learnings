# Day 04 - Pointers in Go

## What I Learned

- What pointers are (& and \*)
- Passing by value vs reference
- Pointer with structs
- Nil pointers and memory behavior
- &x → gives address of x
- \*p → gives value at that address

## 📝 Notes

- Pointers allow modifying original data and improve performance.
- In Go, everything is passed by value, including pointers.
- Go automatically handles &user for you when using pointer receivers.
- Dereferencing a nil pointer causes a panic.
- A pointer that has not been initialized to any memory address is called a nil pointer.
- Use pointer when:
  - You want to modify original data
  - Struct is large (performance)
  - Sharing data across functions
- Don’t use pointer when:
  - Data is small (int, bool)
  - No modification needed

## Practice

- Increment using value vs pointer
- Struct with pointer receiver
- Nil pointer check
