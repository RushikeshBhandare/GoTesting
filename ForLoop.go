package main

import (
	"fmt"
)

func main() {
	x := 1

	// for staament with single condition
	for x <= 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Donn")

	// normal for statement
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Normal For Loop IS Done")

	//Another Way of Fo Loop

	v := 0
	for {
		if v > 9 {
			break
		}
		fmt.Println(v)
		v++
	}
	fmt.Println("Loop Over")

	// continue in for loop
	S := 0
	for {
		S++
		if S > 100 {
			break
		}
		if S%2 != 0 {

			continue
		}
		fmt.Println(S)

	}
}
