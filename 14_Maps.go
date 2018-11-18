package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Ram"] = 40
	grades["Sam"] = 90
	grades["Tom"] = 60

	fmt.Println(grades)

	RamsGrade := grades["Ram"]
	fmt.Println(RamsGrade)

	delete(grades,"Ram")
	fmt.Println(grades)

	for key,value := range grades {
		fmt.Println(key,":",value)
	}


}