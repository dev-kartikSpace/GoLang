//go:build ignore

package main

import "fmt"

func main() {
	type Student struct {
		firstName string
		lastName  string
		age       int
		subject   []string
	}
	var Student1 Student

	//initialisig a struct (implicit)
	Student1 = Student{
		"code",
		"sharma",
		20,
		[]string{"maths", "english"},
	}
	fmt.Printf("%+v\n", Student1)

	//explicit itialisation
	Student2 := Student{
		firstName: "kartik",
		lastName:  "sharma",
	}

	fmt.Printf("%+v\n", Student2)

	fmt.Println(Student1.firstName)

	Student2.subject = append(Student2.subject, "maths", "arts")
	fmt.Printf("%+v\n", Student2)
}
