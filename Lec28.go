//go:build ignore

package main

import "fmt"

func main() {
	age := 14
	if age > 18 {
		fmt.Println("you are an adult")
	} else if age > 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a minor")
	}

	// if even := isEven(4); even {
	// 	fmt.Println("sahi hai")
	// }
}
