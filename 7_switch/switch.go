package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
		//no need to write break
	case 2:
		fmt.Println("two")
	default://no need to write deffault statements
		fmt.Println("other")
	}


	//multiple condition switch

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("workday")
	}

	//Type switch
	whoAmI := func(i interface{}){
		switch t := i.(type) /*or switch i.(type)*/{
		case int: 
			fmt.Println("Integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("golang") //returns string
	whoAmI(50) //returns Integer
	whoAmI(5.5) //returns other
}
