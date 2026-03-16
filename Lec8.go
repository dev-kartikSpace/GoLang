//go:build ignore

package main

import "fmt"

func main() {
	//_ is considered a letter
	var myInt int = 10_000_000
	//dont use bool, break etc .. any pre declared identifiers
	fmt.Println(myInt)

}
