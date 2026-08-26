package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("Person is a kid")
	}

	var role = "admin"
	var hasPermissions = false

	if role == "admin" || hasPermissions == true {
		fmt.Println("yes")
	}

	//We can declare a variable inside if construct
	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager")
	}

	//go does NOT have ternary operators
	//so use normal if else
}