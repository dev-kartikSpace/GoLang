//go:build ignore

package main

import "fmt"

func main() {
	//nill map
	//var nameAge map[string]int

	//variable shorthand declaration
	// nameAge := map[string]int{}

	//make fuction
	// nameAge := make(map[string]int)

	nameAge := map[string]int{
		"John": 30,
		"Jane": 25,
	}

	fmt.Println(len(nameAge))
	fmt.Println(nameAge["John"])

	//coma ok idiom
	gradeMap := map[string]int{
		"foo": 90,
		"bar": 10,
		"x":   0,
	}

	gradeX, ok := gradeMap["x"]

	fmt.Println(gradeX, ok)
}
