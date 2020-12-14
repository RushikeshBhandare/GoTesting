package main

import (
	"fmt"
)

type Persone struct {
	first string
	last  string
	age   int
}

func (x Persone) speak() {
	fmt.Println("NAme Of A Persone :- ", x.first)
	fmt.Println("Age Of A Persone :- ", x.age)

}

func main() {
	PersoneOne := Persone{
		first: "Sunny",
		last:  "Bhandare",
		age:   45,
	}

	personeTwo := Persone{
		first: "Rushikesh",
		last:  "Bhandare",
		age:   21,
	}

	PersoneOne.speak()
	personeTwo.speak()

}
