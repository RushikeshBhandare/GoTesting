package main

import (
	"fmt"
)

func main() {
	for i := 33; i <= 100; i++ {
		a := string(i)
		fmt.Print(i, ": its Character \t")
		fmt.Println(a)
	}
}
