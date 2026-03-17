//go:buuild ignore

package main

import "fmt"

func main() {
	a := 10
	b := 5

	complex1 := complex(10, 15)
	complex2 := complex(5, 10)

	sum := a + b
	fmt.Println("Sum: ", sum)

	firstName := "John"
	lastName := "Doe"

	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	fmt.Printf("Full Name : %s %s\n", firstName, lastName)

	product := a * b
	fmt.Println("Product:", product)

	complexSum := complex1 + complex2
	fmt.Println("Complex Sum: ", complexSum)

	remainder := a % b
	fmt.Println("Remainder: ", remainder)

	a = a + 5
	fmt.Println("Updated value of a: ", a)

	//zeroDevision := a / 0

}
