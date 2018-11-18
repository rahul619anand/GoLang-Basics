package main

import "fmt"

func foo(c chan int, value int) {
	c  <- value * 5
}


func main() {
	fooVal := make(chan int)

	// send and receive on channel is blocking
	// hence even though we use go routine we dont need to wait
	go foo(fooVal, 5)
	go foo(fooVal, 10)

	v1, v2 := <- fooVal, <- fooVal

	fmt.Println(v1, v2)
}