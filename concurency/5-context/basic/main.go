package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------------------")

	ctx1 := context.Background()
	withCancel, cancel := context.WithCancel(ctx1)
	go func() {
		for {
			select {
			case <-withCancel.Done():
				fmt.Println("ctx1: Canceled")
				return
			default:
				fmt.Println("ctx1: Sleeping")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	cancel()

	time.Sleep(time.Second * 1)
	fmt.Println("------------------------")

	ctx2 := context.Background()
	withDeadline, _ := context.WithTimeout(ctx2, time.Second)

	loop := true
	for loop {
		select {
		case <-withDeadline.Done():
			fmt.Println("ctx2: Canceled")
			loop = false
		default:
			fmt.Println("ctx2: Sleeping")
			time.Sleep(time.Second)
		}
	}

	fmt.Println("------------------------")

	ctx3 := context.Background()
	withValue := context.WithValue(ctx3, "someKey", 5)
	fmt.Println("ctx3:", withValue.Value("someKey"))

	fmt.Println("------------------------")

	ctx4 := context.Background()
	withCancel, cancel = context.WithCancel(ctx4)
	withCancel, _ = context.WithCancel(withCancel)
	go func() {
		for {
			select {
			case <-withCancel.Done():
				fmt.Println("withCancel: Canceled")
				return
			default:
				fmt.Println("withCancel: Sleeping")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 1)
}
