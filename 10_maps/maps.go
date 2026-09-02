package main

import "fmt"

//Maps aka hash, objects, dictionary, etc

func main(){
	//---Creating Map---
	m1 := make(map[string]string) //First "string" is datatype of key and second is of value

	//---Setting an element---
	m1["name"] = "golang"
	m1["area"] = "backend"

	//---Printing---
	fmt.Println(m1["name"])
	fmt.Println(m1["name"], m1["area"])
	//IMP: If key doesn't exist then it returns zero
	fmt.Println(m1["hello"])
	
}