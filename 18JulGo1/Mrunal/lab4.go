package main

import "fmt"

func main(){
	var str1 string
	var str2 string
	fmt.Println("Enter first string: ")
	fmt.Scanln(&str1)
	fmt.Println("Enter second string: ")
	fmt.Scanln(&str2)

	length := addString(str1,str2)
	fmt.Println("Sum of length of two string is: ",length)
}

func addString(x string,y string )(int){
	length1 := len(x)
	length2 := len(y)
	return length1+length2
}