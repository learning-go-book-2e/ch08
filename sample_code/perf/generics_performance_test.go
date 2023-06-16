package perf

import "testing"

type Ager interface {
	age() int
}

type person struct {
	Name string
	Age  int
}

func (p person) age() int {
	return p.Age
}

type building struct {
	Address string
	Age     int
}

func (b building) age() int {
	return b.Age
}

func doubleAge(a Ager) int {
	return a.age() * 2
}

func doubleAgeGeneric[T Ager](a T) int {
	return a.age() * 2
}

var blackhole int

func BenchmarkGenericFunc(b *testing.B) {
	p := person{
		Name: "Superman",
		Age:  30,
	}
	for i := 0; i < b.N; i++ {
		blackhole = doubleAgeGeneric(p)
	}
}

func BenchmarkFunc(b *testing.B) {
	p := person{
		Name: "Superman",
		Age:  30,
	}
	for i := 0; i < b.N; i++ {
		blackhole = doubleAge(p)
	}
}
