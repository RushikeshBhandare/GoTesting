package main

import (
	"fmt"
)

func main() {
	fmt.Println("Heloo from main function")

	//Calling Fuunction A() with No Arguments
	A()

	//Calling function PrintName() with One Argument Name

	PrintName("Raju")

	// function adress() retuniong the value as a string
	Adr := Adress("Sangli, Maharashtra, india")

	fmt.Println(Adr)

	//calling func Addition

	Addition(1, 2, 4, 4, 5, 6)

	//Anonymous function
	func() {
		fmt.Println("Printing From Anonymous Function")
	}()

	//Anonymous function with arguments
	func(x int, y int) {
		Sum := x + y
		fmt.Println("Sum of x: - ", x, " And y:- ", y, " IS :- ", Sum)
	}(15, 70)

	//Funciton Expression
	Exp := func() {
		fmt.Println("Calling from Function Expression")
	}

	Exp()
}

//Crreating Function With no Pramiters

func A() {
	fmt.Println("Calling From A Function ()")
}

//Function with Paramitrs
//Everything in go is PassBY Value

func PrintName(Name string) {
	fmt.Println("Name is :- ", Name)
}

//return

func Adress(Ad string) string {
	return fmt.Sprint("Adress IS :- ", Ad)
}

//Variatic paramiter

func Addition(x ...int) {
	//Adding All the numers

	sum := 0
	for i := 0; i < len(x); i = i + 1 {
		sum = sum + x[i]
	}
	fmt.Println("Addition of :- ", x, " IS :- ", sum)

	//Function Returning A Function
	RF := ReturnFunction()

	fmt.Println("Calling Fom REturning Function")
	fmt.Printf("%T", RF)
	fmt.Println()

}

//Returning A Funvction

func ReturnFunction() func() int {
	return func() int {
		return 128
	}
}
