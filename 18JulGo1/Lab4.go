package main

import "fmt"

func main() {

	s := ""
	q := ""
	fmt.Println("Enter the strings")
	fmt.Scanln(s, q)
	sum := len(s) + len(q)
	fmt.Println(sum)

}
