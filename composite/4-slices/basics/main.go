package main

import (
	"fmt"
)

func main() {
	fmt.Println("1--------------------------")

	arr := [5]int{1, 2, 3, 4, 5}

	slice := arr[1:3]
	fmt.Println(slice)

	fmt.Println("2--------------------------")

	fmt.Println(&arr[1])
	fmt.Printf("%p\n", &slice[0])

	fmt.Println("3--------------------------")

	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", slice) // Pointer to the underlying array

	fmt.Println("4--------------------------")

	short := []int{1, 2, 3}
	fmt.Println(short)

	fmt.Println("5--------------------------")

	shortMake := make([]int, 5, 10)
	fmt.Println(shortMake)
}
