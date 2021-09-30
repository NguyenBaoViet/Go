package main

import "fmt"

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	x := truck{
		vehicle: vehicle{
			doors: "2",
			color: "blue",
		},
		fourWheel: true,
	}

	y := sedan{
		vehicle: vehicle{
			doors: "1",
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("truck:")
	fmt.Println("\tDoors:", x.vehicle.doors)
	fmt.Println("\tColor:", x.vehicle.color)
	fmt.Println("\tFour Wheel:", x.fourWheel)

	fmt.Println("sedan:")
	fmt.Println("\tDoors:", y.vehicle.doors)
	fmt.Println("\tColor:", y.vehicle.color)
	fmt.Println("\tLuxury:", y.luxury)
}
