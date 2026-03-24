//go:build ignore

package main

import "fmt"

func main() {
	// var a [5]int
	// fmt.Println(a)

	// var a = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(a)

	//sparse array declaration
	// a := [5]int{0: 1, 3: 5, 6}
	// fmt.Println(a)

	//implicit ;ength declaration
	// a := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(len(a))
	// fmt.Println(a)

	//accessing the elements of the array
	// b := a[0]
	// fmt.Println(b)

	//2D array
	// a := [2][3]int{{1, 2}, {3, 4, 5}}
	// fmt.Println(a)

	a := [][]int{{1, 2}, {3, 4, 5}}
	fmt.Println(a)
}
