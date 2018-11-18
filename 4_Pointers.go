package main

import "fmt"

func main() {
	a := 10
	b := &a 
	fmt.Println(b) // memory address
	fmt.Println(*b) // value at the memory address

	*b = 5

	fmt.Println(a)
	fmt.Println(*b)

	*b = *b**b

	fmt.Println(*b)


}