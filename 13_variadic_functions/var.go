package main

import "fmt"


//Variadic Functions means You can pass n number of
//  arguments to the function like fmt.Println()

func sum(nums ...int) int {
	total := 0

	for _, num:= range nums{
		total = total + num
	}
	return total
}




func main() {
	fmt.Println(1, 2)

	r := sum(1,2,3,4)
	fmt.Println(r) //10

	//if you have a slice
	nums := []int{3,4}
	result := sum(nums...)
	fmt.Println(result) //7
}
