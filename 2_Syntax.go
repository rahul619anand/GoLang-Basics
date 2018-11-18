package main	

import ("fmt"
		"math/rand"
		"math")

func random() {
	fmt.Println (" The random number is ",rand.Intn(100))
}

func main() {
	fmt.Println (" The square root of 4 is ", math.Sqrt(4))
	random()
}



// TO check docs :   godoc math/rand Intn