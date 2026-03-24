package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:]

	b = append(b, 15)
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}
