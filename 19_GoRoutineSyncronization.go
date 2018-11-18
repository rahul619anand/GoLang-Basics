package main

import ("fmt"
		"time"
		"sync")

var wg sync.WaitGroup 


func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond*100)
		fmt.Println(s)
	}
	wg.Done()
	
}

func main() {
	wg.Add(1)
	go say("hi")
	wg.Add(1)
	go say("rahul")
	wg.Wait()

	wg.Add(1)
	go say("bye")
	wg.Wait()
}