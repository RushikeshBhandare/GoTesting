package main

import "fmt"

//Struct persone to store the value of name and number at a smae place
type persone struct {
	Name   string
	Number string
}

func main() {

	//create the arry of struct persone
	p1 := [4]persone{}

	// assign the values to the struct of array p1
	p1[0].Name = "Sunny"
	p1[0].Number = "555-666-0100"

	p1[1].Name = "Raju"
	p1[1].Number = "555-666-0102"

	p1[2].Name = "Golu"
	p1[2].Number = "555-666-0103"

	p1[0].Name = "Damu"
	p1[1].Number = "555-666-0104"

	//looping through all the array to find given name
	for i := 0; i < 4; i++ {
		//checking if the name is in the given array or not
		if p1[i].Name == "Golu" {
			fmt.Println(p1[i].Number)
			return
		}
	}

	fmt.Println("Name Not found")

}
