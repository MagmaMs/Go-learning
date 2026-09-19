package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}

	//here "_" means index
	for _, num := range nums {
		fmt.Println(num) //6 7 8
	}

	sum := 0

	for _, num := range nums {
		sum += num
		fmt.Println(sum) //21
	}

	for i, num := range nums {
		fmt.Println(num, i) //6 0, 7 1, 8 2
	}

	m := map[string]string{"name": "Go"}

	for k, v := range m {
		fmt.Println(k, v) //name Go
	}

	for k := range m {
		fmt.Println(k) //name (only keys)
	}

	for v := range m {
		fmt.Println(v) //Go (only values)
	}

	//unicode code for each letter
	//unicode code point rune
	//starting byte of rune
	//255 -> 1 byte, 255+ -> 2 byte
	for i, c := range "golang" {
		fmt.Println(i, c)         //0 103, 1 111, 2 108, 3 97, 4 110, 5 103
		fmt.Println(i, string(c)) // g o l a n g
	}
}
