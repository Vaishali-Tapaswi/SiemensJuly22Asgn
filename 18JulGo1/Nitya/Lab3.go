package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Enter a numbers")
	fmt.Scanln(x)
	calc(x)
	fmt.Println(x)

}

func calc(x int) (int, int) {
	return x * x, x * x * x
}
