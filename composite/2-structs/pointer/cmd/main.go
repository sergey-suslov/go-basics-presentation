package main

import (
	"fmt"
)

func mutate(a Human, b *Human) {
	a.age = 100
	b.age = 100
}

// Ordered in memory
type Human struct {
	name string
	age  uint
}

func main() {
	fmt.Println("--------------------------")

	a := Human{"A", 10}
	b := &Human{"B", 10}
	mutate(a, b)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("--------------------------")

	fmt.Println((*b).name)
	fmt.Println(b.name)

	fmt.Println("--------------------------")

	empty := new(Human) // ~ &Human{}
	fmt.Println(empty)
}
