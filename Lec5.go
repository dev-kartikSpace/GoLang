//go:build ignore

package main

import "fmt"

func main() {
	var myString string
	fmt.Println(myString)

	myString = "Hello Kartik"
	fmt.Println(myString)

	// \n is to go to a new line
	myString = "Hellowww\nworld"
	fmt.Println(myString)

	//lines and tabs can be added using backticks
	myString = `Hello
	world`
	fmt.Println(myString)

	var firstName, lastName string
	firstName = "kartik"
	lastName = "sharma"
	//var fullName = firstName + " " + lastName
	//var fullName = fmt.Sprintf("%s , %s\n", firstName , lastName)
	// fmt.Println(fullName)
	fmt.Printf("%s %s\n", firstName, lastName)

}
