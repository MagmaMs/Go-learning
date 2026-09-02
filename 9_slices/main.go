package main

import "fmt"

//most used construct in go
func main(){

	//Uninitialised slice is Nil
	var nums []int
	fmt.Println(nums)
	//prints "[]"

	fmt.Println(len(nums))
	//prints "0" because it is nil

}