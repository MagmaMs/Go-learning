package main

import (
	"fmt"
	"slices"
)

// most used construct in go
func main() {

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
	fmt.Println(nums2)      //Prints [0 0 1]
	fmt.Println(cap(nums2)) //prints 5

	//if we append more
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	fmt.Println(cap(nums2)) //prints 10
	fmt.Println(nums2)      //Prints [0 0 1 2 3 4]

	//Thus second argument of "make" initialises number of zeroes.

	//Another way to initialise Slice

	num := []int{} //num := []int{1,2,3}
	fmt.Println(num)
	fmt.Println(cap(num))
	fmt.Println(len(num))
	//Prints [] 0 0
	num = append(num, 1)
	//Prints [1] 1 1

	//Using index
	num[0] = 3 //Might give error if index is more than the length entered in make
	//Prints [3 0] if length is 2, first we entered and second is zero by default

	//-----Copy Function------
	var copy1 = make([]int, 0, 5)
	copy1 = append(copy1, 2)
	var copy2 = make([]int, len(copy1))

	//copy1 = append(copy1, 2)
	fmt.Println(copy1, copy2) //Prints [2] []

	copy(copy2, copy1) //we need to append first then initialise second slice using "len(copy1)"
	//now prints [2] [2]

	//-----Slicing indices-----

	xyz := []int{1, 2, 3, 4, 5}
	fmt.Println(xyz[0:2]) //excludes last index i.e. 2
	//prints [1 2]

	//---Slices package---
	var abc = []int{1, 2}
	var def = []int{1, 2}
	fmt.Println(slices.Equal(abc, def)) //returns true

	//---2D Slices---
	var twoD = [][]int{{1, 2}, {3, 4}} //prints [[1 2] [3 4]]
	fmt.Println(twoD)
}
