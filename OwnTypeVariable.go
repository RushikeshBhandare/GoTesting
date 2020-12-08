package main

import (
	"fmt"
	"runtime"
)

type Dg int

var x Dg

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)

}
