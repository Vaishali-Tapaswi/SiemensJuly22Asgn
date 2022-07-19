package main

import "fmt"

func main() {

	name := ""
	fmt.Println("Enter your Name:")
	fmt.Scanln(&name)

	age := 0
	fmt.Println("Enter you age")
	fmt.Scanln(&age)

	licenseType := ""
	fmt.Println("What type of license do you want: two-wheeler or four-wheeler")
	fmt.Scanln(&licenseType)

	switch licenseType {
	case "two-wheeler":
		if age > 16 {
			fmt.Println("Eligible")
		} else {
			fmt.Println("Not Eligible")
		}
	case "four-wheeler":
		if age > 18 {
			fmt.Println("Eligible")
		} else {
			fmt.Println("Not Eligible")
		}
	}

}
