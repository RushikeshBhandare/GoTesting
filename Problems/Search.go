package main

import (
	"fmt"
)

func main() {
	Names := []string{"sunny", "Rushiokesh", "Bhandare"}

	fmt.Println(Names)

	for i := 0; i < 2; i++ {
		if Names[i] == "Sunny" {
			fmt.Println("Found :- " + Names[i])
			return
		}
	}
	fmt.Println("Sunny Not Found")

}
