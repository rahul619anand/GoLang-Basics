package main


import "fmt"


func add (a,b float64) float64 {
	return a + b
}


func tuple (a,b string) (string,string) {
	return a,b
}

const num3 float64 = 6.1

func main() {

	var num1,num2 float64 = 5.6, 5.5
	fmt.Println("Add result", add(num1, num2))

	// here type cannot be changed after compilation
	num4 := 5.5
	fmt.Println("Again Add result", add(num3, num4))

	w1,w2:="Hey","there"
	fmt.Println(tuple(w1, w2))

	/*
	Multi line
	comment
	*/
	
	// type inference
	var a int = 64
	var b float64 = float64(a)
	
	// x will be type int
	x:=a 

}