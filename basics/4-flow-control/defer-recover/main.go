package main

import "fmt"

func level1() {
	defer func() {
		fmt.Println("defer - 1")
		if err := recover(); err != nil {
			fmt.Println("level 1 err:", err)
		}
	}()
	level2()
	fmt.Println("level 1")
}

func level2() {
	defer func() {
		fmt.Println("defer - 2")
		if err := recover(); err != nil {
			fmt.Println("level 2 err:", err)
		}
	}()
	level3()
	fmt.Println("level 2")
}

func level3() {
	defer func() {
		fmt.Println("defer - 3")
	}()

	panic("panic message")
	fmt.Println("level 3")
}

func main() {

	level1()
}
