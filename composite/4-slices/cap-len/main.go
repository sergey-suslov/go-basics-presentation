package main

import (
	"fmt"
)

func printSliceAddress(slice []int) {
	fmt.Printf("Slice address in printSliceAddress = %p\n", slice)
}

func main() {
	fmt.Println("--------------------------")

	var defaultSlice []int
	fmt.Println(defaultSlice == nil)
	fmt.Println(len(defaultSlice))
	fmt.Println(cap(defaultSlice))

	fmt.Println("--------------------------")

	short := []int{1, 2, 3}
	fmt.Printf("Slice address = %p\n", short)
	printSliceAddress(short)
	fmt.Println(len(short))
	fmt.Println(cap(short))

	fmt.Println("--------------------------")

	shortMake := make([]int, 5, 10)
	fmt.Println(len(shortMake))
	fmt.Println(cap(shortMake))

	fmt.Println("--------------------------")

	slice := make([]int, 1, 2)
	fmt.Printf("len = %d\n", len(slice))
	fmt.Printf("cap = %d\n", cap(slice))
	fmt.Printf("pointer = %p\n", slice)

	// Grows the slice; Updates pointers to the slice
	for len(slice) < 5 {
		fmt.Printf("\nAppending...\n")
		slice = append(slice, 3)
		fmt.Printf("len = %d\n", len(slice))
		fmt.Printf("cap = %d\n", cap(slice))
		fmt.Printf("pointer = %p\n", slice)
	}

	fmt.Println("--------------------------")

	for i, v := range slice {
		fmt.Printf("[%d]=%d ", i, v)
	}
}
