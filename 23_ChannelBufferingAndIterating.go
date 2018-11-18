package main

import ("fmt"
		"sync")

var wg sync.WaitGroup 

func foo(c chan int, value int) {
	defer wg.Done()
	c  <- value * 5
}


func main() {
	fooVal := make(chan int, 10) // buffer for 10 items

	for i := 0; i < 10 ; i++ {
		wg.Add(1)
		go foo(fooVal,i)
	}

	wg.Wait()
	close(fooVal) // close channel

	for item := range fooVal {
		fmt.Println(item)
	}
}