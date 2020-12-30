package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 40

	}( )

	fmt.Println(<-ch)

}
