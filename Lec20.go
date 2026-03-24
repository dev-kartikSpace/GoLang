//go:build ignore

package main

import "fmt"

func main() {
	// var a []int
	//nil slice : declared but not initialized
	// fmt.Println(a == nil)

	// a := []int{1, 2, 3, 4, 5}
	// fmt.Println(a)

	// a := []int{5, 2: 10, 50}
	// fmt.Println(a)

	//make function
	// a := make([]int, 5)
	// fmt.Println(a)
	// fmt.Println(cap(a))

	// a := make([]string, 5, 10)
	// fmt.Println(cap(a))

	//Append function
	a := make([]int, 5)
	a = append(a, 10)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(len(a))

	someFunction(a)
}

func someFunction(x []int) {}
