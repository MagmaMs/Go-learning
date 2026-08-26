package main

import "fmt"

const roll = 5

func main() {
	const name = "golang"
	const age = 30
	fmt.Println(name)
	fmt.Println(roll)

	//Multiple constants together
	const (
		a = 50
		b = "abc"
	)

	fmt.Println(a, b)
}
