
package main

import ("fmt"
		"time"
		"sync")

var wg sync.WaitGroup 


func cleanup() {
	// recover the panic 
	if r := recover(); r != nil {
		fmt.Println(" Recovered in cleanup", r)
	}
	defer wg.Done() // last call before func exits
	// returns to WaitGroup
}

func say(s string) {
	
	defer cleanup() // last call before func exits
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond*100)
		fmt.Println(s)
		if i == 2 {
			panic("oh no, got 2") // exception 
		}
	}
	
}

func main() {
	wg.Add(1)
	go say("hi")
	wg.Wait()
}