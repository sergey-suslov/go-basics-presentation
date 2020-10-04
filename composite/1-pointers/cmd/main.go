package main

import "fmt"

func main() {
	var pointer *int
	fmt.Println(pointer)

	fmt.Println("--------------------------")

	i, j := 42, 2701

	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	fmt.Println("--------------------------")

	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(j * 37)
}
