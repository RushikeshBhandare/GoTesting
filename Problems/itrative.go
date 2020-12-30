package main

import "fmt"

func main() {

	var b int
	fmt.Scanln(&b)

	Draw(b)
}

func Draw(s int) {
	for i := 0; i <= s; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
