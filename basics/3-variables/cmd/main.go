package main

import "fmt"

func main() {
	//bool
	//
	//string
	//
	//int  int8  int16  int32  int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//
	//byte // alias for uint8
	//
	//rune // alias for int32
	//// represents a Unicode code point
	//
	//float32 float64
	//
	//complex64 complex128

	var numeric int
	var boolean bool
	var string string

	fmt.Printf("numeric = %d\n", numeric)
	fmt.Printf("boolean = %t\n", boolean)
	fmt.Printf("string = \"%s\"\n", string)
	fmt.Println("------------------------")

	var num int = 1
	var float = 1.0
	short := 2.0

	fmt.Println(num)
	fmt.Println(float)
	fmt.Println(short)
	fmt.Println("------------------------")


	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Println(u)
	fmt.Println("------------------------")

	fmt.Println(Pi)
}

const Pi = 3.14
