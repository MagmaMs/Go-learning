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

	//---Delete---
	delete(m1, "area")
	fmt.Println(m1)
	//prints "map[name:golang]"

	//---Clear all---
	clear(m1)

	//---Another way of Initialisation
	m2 := map[int]string{1:"one", 2:"two"}
	fmt.Println(m2)

	//
	_, ok := m2[1]

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not okay")
	}
}