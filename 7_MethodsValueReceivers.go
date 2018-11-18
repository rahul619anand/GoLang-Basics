package main 

import "fmt"

// 2 types of methods
// 1. value receivers
// 2. pointer recievers


const u16bitmax float64 = 65535
const kmh_multiple float64 = 1.60934


// 1. value receivers
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/u16bitmax)
}

// 1. value receivers
// func associated to a struct (car here)
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/u16bitmax/kmh_multiple)
}

type car struct {
	gas_pedal uint16 // min 0 max 65535
	brake_pedal uint16
	steering_wheel int16  // -32k to + 32k
	top_speed_kmh float64
}

func main() {
	a_car := car{gas_pedal: 65000,
				 brake_pedal: 0, 
				 steering_wheel: 12561, 
				 top_speed_kmh: 225.0}

	fmt.Println(a_car.top_speed_kmh)	

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())		 

}
