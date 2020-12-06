package main

import (
	"fmt"
)

var x int = 42
var y string = "James bond"
var z bool = true

func main() {

	S := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(S)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}
