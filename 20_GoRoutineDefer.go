package main

import ("fmt"
		"time"
		"sync")

var wg sync.WaitGroup 


func say(s string) {
	defer wg.Done() // also runs when panic happens 
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond*100)
		fmt.Println(i)
		fmt.Println(s)
	}
	
}

func main() {
	wg.Add(1)
	go say("hi")
	wg.Wait()
}