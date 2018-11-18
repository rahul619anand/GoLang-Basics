package main

import "fmt"

// gofmt filename.go - formats file
// go has structs as it doesn't has classes

type car struct {
	gas_pedal uint16 // min 0 max 65535
	brake_pedal uint16
	steering_wheel int16  // -32k to + 32k
	top_speed_kmh float64
}

func main() {
	a_car := car{gas_pedal: 10,
				 brake_pedal: 0, 
				 steering_wheel: 5, 
				 top_speed_kmh: 100}
	
	// short form
	//a_car := car{10,0,5,100}

	fmt.Println(a_car.top_speed_kmh)			 

}