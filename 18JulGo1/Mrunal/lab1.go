package main

import "fmt"

func main() {

    var name string
	var age int
	var vehicle int

	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("Enter your age")
	fmt.Scanln(&age)
	fmt.Println("Enter your choice : 1)two wheeler 2)four wheeler")
	fmt.Scanln(&vehicle)

	if age > 16 && age < 18 {
		fmt.Println("You are elgible for Only two wheeler")
	} else if 18 < age && age < 70 {
		fmt.Println("You are elgible for two wheeler or four wheeler")
	} else {
		fmt.Println("Invalid age for license")
	}



}