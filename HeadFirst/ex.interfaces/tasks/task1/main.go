package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("direction", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeeding up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {

	var vehicle Vehicle = Car("Toyota Yarvic")
	vehicle.Accelerate()
	vehicle.Brake()
	vehicle.Steer("left")
	fmt.Println(vehicle)

	vehicle = Truck("Fnord F80")
	vehicle.Accelerate()
	vehicle.Brake()
	vehicle.Steer("right")
}
