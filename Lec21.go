//go:build ignore

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	//b := a[2:]
	// b = append(b, 10)
	// fmt.Println(a)
	// fmt.Println(b)

	//passig capacity to sliced array
	b := make([]int, 6)
	copy(b, a[2:])
	fmt.Println(b)

}
