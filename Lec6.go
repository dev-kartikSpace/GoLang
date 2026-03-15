//go:build ignore

package main

import "fmt"

//var age int
//var age int = 10
//var age = 10
//var age, name = 10, "hello world"
//age := 10
//age , name := 10, "kartik sharma"

// var myInt int = 10

var (
	myInt    int    = 10
	myString string = "Hello world"
)

func main() {
	fmt.Println(myInt)
	fmt.Println(myString)
	//something()

	age := 30
	age = 50
	fmt.Println(age)

	var test int = 10
	println(test)
}

// func something() {
// 	fmt.Println("something:", myInt)
// }
