package main

import "fmt"

func main() {
	var nums [5]int

	fmt.Println(len(nums)) //returns array length

	nums[0] = 1
	fmt.Println(nums[0]) //returns 1
	fmt.Println(nums) //returns [1 0 0 0] because 0 is default value for int, false is default for bool and empty string for strings

	nums1 := [3]int{1,2,3}
	fmt.Println(nums1)
}
