package main

import "fmt"

func main() {
	var smallPositiveValue uint8
	smallPositiveValue = 10
	fmt.Println(smallPositiveValue)

	var smallPosNegInt int8
	smallPosNegInt = -10
	fmt.Println((smallPosNegInt))

	var myInt int = 23000000043
	fmt.Println(myInt)

	myInt = int(smallPositiveValue)
	myInt = int(smallPosNegInt)
	fmt.Println(myInt)

	var MyByte = 'A'
	fmt.Println(MyByte)
	//byte is used to represent raw data
	//byte also distinguish between int and uint
	//byte is an alias for uint8

	//rune is alias for int32
	var myRune = 'B'
	println(myRune)

}
