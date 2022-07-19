package main

import (

	"fmt"
	"os"
	"strconv"
)

func main() {
	number := os.Args[1]
	i, _ := strconv.Atoi(number)
	printFibonacciSeries(i)
}

func printFibonacciSeries(num int) {
	a := 0
	b := 1
	c := b

    fmt.Printf("Series is: %d %d", a, b)
	for true {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}

}