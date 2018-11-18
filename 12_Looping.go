package main 

import "fmt"

func main() {

	for i:= 0; i < 2; i++ {
		fmt.Println(i)
	}

	j:=0
	for j <10 {
		fmt.Println(j)
		j+=2
	}

	x := 5
	for {
		fmt.Println("hi")
		x*=6
		if x > 35 {
			break
		}
	}

	// for {
	// 	fmt.Println("infinite loop")
	// }
	
}