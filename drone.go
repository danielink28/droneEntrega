package main

import (
	"domicilio/orientation"
	"fmt"
)

type Drone struct {
	X           int
	Y           int
	Orientation orientation.Orientation
}

func NewDrone() (drone Drone) {
	drone = Drone{}
	drone.Orientation = orientation.North{}
	return
}

func (drone *Drone) move() {
	drone.X, drone.Y = drone.Orientation.Move(drone.X, drone.Y)
}

func (drone *Drone) rotate(giro string) {
	drone.Orientation = drone.Orientation.Rotate(giro)
}

func (drone *Drone) position() (position string) {
	position = fmt.Sprintf("(%d,%d) direccion %s", drone.X, drone.Y, drone.Orientation.Position())
	return
}