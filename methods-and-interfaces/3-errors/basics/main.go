package main

import (
	"errors"
	"fmt"
)

var CustomError = errors.New("custom error")

func doWork() (string, error) {
	return "work is done", CustomError
}

func traceError() (out string, err error) {
	defer func() {
		if err != nil {
			fmt.Println("Error traced:", err)
		}
		fmt.Println("Out traced:", out)
	}()

	return "output", CustomError
}

func main() {
	fmt.Println("--------------------------")

	_, err := doWork()
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	fmt.Println("--------------------------")

	_, err = traceError()
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
}
