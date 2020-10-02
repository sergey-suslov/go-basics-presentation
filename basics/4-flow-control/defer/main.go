package main

import (
	"fmt"
	"time"
)

func deferFunc() {
	defer func() {
		fmt.Println("deferFunc")
	}()
}

func deferWithParamFunc() {
	defer func(num int) {
		fmt.Println("deferWithParamFunc")
	}(printAndReturn(1))
	time.Sleep(2 * time.Second)
}

func multiDeferFunc() {
	defer func() {
		fmt.Println("multiDeferFunc 1")
	}()
	defer func() {
		fmt.Println("multiDeferFunc 2")
	}()
}

func deferFuncWithResult() (result int) {
	defer func() {
		fmt.Printf("deferFuncWithResult %d", result)
	}()

	return 1
}

func main() {
	deferFunc()
	deferWithParamFunc()
	multiDeferFunc()
	_ = deferFuncWithResult()
}

func printAndReturn(num int) int {
	fmt.Printf("printAndReturn -> %d\n", num)
	return num
}
