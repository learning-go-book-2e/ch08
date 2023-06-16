package main

import "fmt"

// Write a generic function that doubles the value of any integer
// or float that's passed in to it. Define any needed generic interfaces.

type ValidTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Doubler[T ValidTypes](t T) T {
	return t * 2
}

func main() {
	fmt.Println(Doubler(10))
	fmt.Println(Doubler(11.2))
}
