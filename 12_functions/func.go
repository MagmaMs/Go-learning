package main

import "fmt"

//func-keyword function-name (argument datatype) return-datatype
func sum(a int, b int) int { //Or sum(a, b int) this puts int for all arguments
	return a + b
}

func getLanguages() (string, string, string) {
	return "a", "b", "c"
}

func diffreturntype() (int, bool, string){
	return 32, true, "why"
}

//Takes function as an argument
func processIt(fn func(a int) int){
	fn(1)
}

//Returns Function
func processIt2() func(a int) int{
	return func(a int) int{
		return 4
	}
}

func main(){
	result := sum(3, 5)
	fmt.Println(result) //8

	l1, l2, l3 := getLanguages()
	fmt.Println(l1,l2,l3)//a b c

	m1, m2, m3 := diffreturntype()
	fmt.Println(m1,m2,m3)//32 true why

	fn := func(a int)int{ //Anonymous function
		return 2
	}
	processIt(fn)
}
