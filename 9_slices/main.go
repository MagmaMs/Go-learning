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

	//--------------------------------

	var nums1 = make([]int, 2)
	fmt.Println(nums1) //prints "[0 0]"
	//first argument is datatype
	//second argument is initial length
	//third argument is initial capacity, if not mentioned then initial capacity = length
	fmt.Println(cap(nums))
	//prints "2"
	//Cap -> Capacity -> Maximum numbers of elements can fit

	var nums2 = make([]int, 2, 5)
	nums2 = append(nums2, 1)
	fmt.Println(nums2) //Prints [0 0 1]
	fmt.Println(cap(nums2)) //prints 5

	//if we append more
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	fmt.Println(cap(nums2))//prints 10
	fmt.Println(nums2) //Prints [0 0 1 2 3 4]

	//Thus second argument of "make" initialises number of zeroes.
}