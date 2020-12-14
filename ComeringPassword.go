package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	UserPassword := "Password1234"

	bs, err := bcrypt.GenerateFromPassword([]byte(UserPassword), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(UserPassword)
	fmt.Println(string(bs))

	loginPassword := "Password1234"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println(err)
		fmt.Println("------can log in-------")
		fmt.Println("PAssword Does not match")
	} else {
		fmt.Println("--------Can log in --------")
		fmt.Println("User can Login ")
	}
	/*
		fmt.Println("----User login password----")
		fmt.Println(loginPassword)
		fmt.Println(string(Up))

		if string(bs) == string(Up) {

			fmt.Println("--------Can log in --------")
			fmt.Println("User can Login ")
		} else {
			fmt.Println("------can log in-------")
			fmt.Println("PAssword Does not match")
		}
	*/
}
