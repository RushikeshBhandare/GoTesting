package main

import "fmt"

type Persone struct {
	Name    string
	age     int
	address string
	pincode int
}

type College struct {
	Persone     Persone
	collageName string
	Cousrse     string
}

func main() {
	p1 := College{
		Persone: Persone{
			Name:    "Sunny Bhandare",
			age:     21,
			address: "Dhamani",
			pincode: 416314,
		},
		collageName: "PDVP Collage Tasgaon",
		Cousrse:     "BCA",
	}
	fmt.Println(p1)
}
