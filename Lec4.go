//go:build ignore

package main

import "fmt"

func main() {
	var smallFloat float32
	fmt.Println(smallFloat)
	smallFloat = 23.098768769876
	fmt.Println(smallFloat)
	//output is 23.098768

	var bigFloat float64
	fmt.Println(bigFloat)
	bigFloat = 23.09876543217654567
	fmt.Println(bigFloat)
	//output is 23.098765432176545

	//complex numbers
	var myComplex complex128
	myComplex = complex(bigFloat, bigFloat)
	fmt.Println(myComplex)

	//to seperate real and imaginary part of complex number
	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)
	fmt.Println(myRealPart, myImaginaryPart)
}
