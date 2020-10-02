package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum > 0 {
		sum -= 2
	}
	fmt.Println(sum)

	for {
		fmt.Println("Infinite cycle")
		time.Sleep(3 * time.Second)
	}
}

