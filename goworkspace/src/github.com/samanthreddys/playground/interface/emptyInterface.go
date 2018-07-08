package main

import "fmt"

//Vehicle interface
type Vehicles interface {
}

//Vehicle struct
type Vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

//Boat str
type Boat struct {
	Vehicle
	length int32
}

// Car Struct
type Car struct {
	Vehicle
	Seats int
}

//Plane struct
type Plane struct {
	Vehicle
	brand string
}

func main() {

	Prius := Car{Vehicle{12, 2, "Black"}, 32}
	MaxBoat := Boat{}
	Emirates := Plane{}

	rides := []Vehicles{Prius, MaxBoat, Emirates}
	for Key, Value := range rides {
		fmt.Println(Key, ": -", Value)
	}
}
