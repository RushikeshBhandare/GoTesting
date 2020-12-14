package main

import "fmt"

type Vehicle struct {
	Door  int
	color string
}

type truck struct {
	Vehicle
	FourWheel bool
	Luxury    bool
}

type sedan struct {
	Vehicle
	FourWheel bool
	Luxury    bool
}

func main() {
	Truck := truck{
		Vehicle: Vehicle{
			Door:  4,
			color: "Black",
		},
		FourWheel: true,
		Luxury:    false,
	}
	Sedan := sedan{
		Vehicle: Vehicle{
			Door:  2,
			color: "Red",
		},
		FourWheel: false,
		Luxury:    true,
	}
	fmt.Println(Truck, Sedan)

	fmt.Println("Tuck Color :- ", Truck.color)
	fmt.Println("Sedan Colr :-", Sedan.color)
}
