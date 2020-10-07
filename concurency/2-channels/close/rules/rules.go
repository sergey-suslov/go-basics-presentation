package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------------------")

	c := make(chan int)

	close(c)

	// Reading from closed channel returns zero value
	v, ok := <-c
	fmt.Println(v, ok)

	// Writing to closed channel triggers panic
	c <- 5

	fmt.Println("------------------------")
}
