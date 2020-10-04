package main

import (
	"fmt"
)

func main() {
	fmt.Println("--------------------------")

	arr := [5]int{}
	fmt.Println(arr)

	fmt.Println("--------------------------")

	arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(arr[2])

	fmt.Println("--------------------------")

	arr[0] = 10
	fmt.Println(arr)

	fmt.Println("--------------------------")

	fmt.Printf("%p\n", &arr)
	fmt.Println(&arr[0])

	fmt.Println("--------------------------")
}
