package main

import (
	"fmt"
	"time"
)

func doRequestWithResult(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Request result"
}

func doRequest(c chan struct{}) {
	time.Sleep(3 * time.Second)
	c <- struct{}{}
}

func main() {

	c := make(chan string)
	go doRequestWithResult(c)
	result := <-c
	fmt.Println(result)

	ch := make(chan struct{})
	go doRequest(ch)
	<-ch
	fmt.Println("Done")
}
