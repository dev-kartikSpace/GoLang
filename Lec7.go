//go:build ignore

package main

import "fmt"

// const x = 10       //untyped constant x with value 10
// const y int32 = 20 //typed constant y with type int32 and value 20

const (
	x                     = 10
	y               int32 = 15
	applicationName       = "Lesson 7"
	//values that can be used as constants
	isRunning = false
	character = 'a'
	//complex real, imag, len
)

func main() {
	var a int
	a = x
	fmt.Println(a)

	var b int
	b = int(y) //declaration of b as int and assigning the value of y to b after converting it to int
	fmt.Println(b)

	var c int32
	c = y
	fmt.Println(c)

	// const z = a + int(b) + int(c) cant do it because right side needs to be constatnt
}
