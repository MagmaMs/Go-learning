package main

import "fmt"

//ONLY FOR LOOP IS THERE
func main(){

	//Classic for loop
	for j := 0; j<3; j++{
		// break
		// continue
		fmt.Println(j)
	}

	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	//infinite loop
	// for{
	// 	println("1")
	// }

	//Range

	for k:= range 3 {
		fmt.Println(k)
		//prints 0,1,2
	}
}