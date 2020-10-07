package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------------------")

	a := make(chan int, 1)
	b := make(chan int, 1)

	select {
	case <-a:
		fmt.Println("A")
	case <-b:
		fmt.Println("B")
	default:
		fmt.Println("Default")
	}
}
