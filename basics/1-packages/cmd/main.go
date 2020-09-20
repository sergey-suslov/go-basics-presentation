package main

import (
	"fmt"
	"go-basics-presentation/basics/1-packages/another"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	another.PublicFunction()
	// another.privateFunction()

	another.UsingPrivateFunction()
}
