package main

import (
	"fmt"
	"testing"
)

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c Cat) Quack() {
	fmt.Println(c.Name + " meow")
}

// go:noinline
func BenchmarkTestInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var c Duck = Cat{"grooming"}
		c.Quack()
	}
}
