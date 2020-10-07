package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

	fmt.Println("------------------------")

	a := make(chan int, 1)
	b := make(chan int, 1)

	a <- 5
	b <- 6

	select {
	case <-a:
		fmt.Println("A")
	case <-b:
		fmt.Println("B")
	}
}
