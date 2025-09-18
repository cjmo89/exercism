package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

func (car Car) DisplayDistance() string {
	var s string
	s = fmt.Sprintf("Driven %d meters", car.distance)
	return s
}

func (car Car) DisplayBattery() string {
	var s string
	s = fmt.Sprintf("Battery at %d", car.battery)
	s += "%"
	return s
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	runs := trackDistance / car.speed
	runs += trackDistance % car.speed
	batteryNeeded := runs * car.batteryDrain
	return car.battery >= batteryNeeded
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
