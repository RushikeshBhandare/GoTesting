package main

import "fmt"

func main() {
	Names := [4]string{"Sunny", "Raju", "Gulu", "Babu"}
	Numbers := [4]string{"654-579-0101", "654-579-0102", "654-579-0103", "654-579-0104"}

	for i := 0; i < 4; i++ {
		if Names[i] == "Raju" {
			fmt.Println(Names[i] + " Has :- " + Numbers[i] + " Phone number")
			return
		}
	}

}
