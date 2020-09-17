package main

import "fmt"

func main() {
	drone := NewDrone()
	drone.move()
	drone.move()
	drone.move()
	drone.move()
	drone.rotate("I")
	drone.move()
	drone.move()
	drone.rotate("D")
	fmt.Println(drone.position())

	drone.rotate("D")
	drone.rotate("D")
	drone.move()
	drone.rotate("I")
	drone.move()
	drone.rotate("D")
	fmt.Println(drone.position())

	drone.move()
	drone.move()
	drone.rotate("I")
	drone.move()
	drone.rotate("D")
	drone.move()
	drone.rotate("D")
	fmt.Println(drone.position())
}
