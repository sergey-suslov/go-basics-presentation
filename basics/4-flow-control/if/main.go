package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := 0
	if sum == 0 {
		fmt.Println("True")
	} else if sum == 1 {
		fmt.Println("False")
	} else {
		fmt.Println("False")
	}

	if num := rand.Intn(2); num == 1 {
		fmt.Println("num is 1")
	}
}

