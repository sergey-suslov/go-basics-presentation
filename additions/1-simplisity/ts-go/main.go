package main

import "fmt"

type number struct {
	val int
}

func main() {
	num := []number{{1}, {2}, {3}, {4}, {5}}

	for i := 0; i < len(num); i++ {
		num[i].val += 1
	}
	fmt.Println(num)
}
