# Day 06 - Interfaces in Go

## What I Learned

- Interfaces define behavior
- Implicit implementation in Go
- Polymorphism using interfaces
- Type assertion and type switch

## 📝 Notes

- Go interfaces are simple but powerful and enable flexible system design.
- if a type has all the methods of an interface, it implements that interface implicitly.
- Empty interface can hold any type (like any in ts)
- Type assertion is used to get the value of an interface
- Type switch is used to check the type of an interface

eg. Code Snippets

- Declaring interface

```go
type Speaker interface {
	Speak() string
}
```

- Implementing interface

```go
type Dog struct{}

func (d Dog) Speak() string {
	return "woof!"
}
```

- Using interface

```go
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}
```

- Type assertion

```go
var y interface{} = "hello"
str := y.(string) // type assertion
fmt.Println(str)
```

- Type switch

```go
func checkType(x interface{}) {
	switch y := x.(type) {
	case int:
		fmt.Println("it is int", y)
	case string:
		fmt.Println("it is string", y)
	case bool:
		fmt.Println("it is bool", y)
	case float64:
		fmt.Println("it is float64", y)
	default:
		fmt.Println("it is unknown type")
	}
}
```

## Practice

- Payment processor system
- Type assertion example
