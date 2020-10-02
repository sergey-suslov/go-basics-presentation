package main

import (
	"fmt"
	"go-basics-presentation/basics/2-functions/function"
	"go-basics-presentation/basics/2-functions/multireturn"
)

func main() {
	fmt.Printf("1 + 2 = %d \n", function.Add(1, 2))
	fmt.Printf("2 + 4 = %d \n", function.AddFancy(2, 4))

	sum, multi := multireturn.ReturnPair(2, 4)
	fmt.Printf("2 + 4 = %d, 2 * 4 = %d \n", sum, multi)
}
