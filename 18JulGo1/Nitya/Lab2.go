package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter length:")
	count, _ := strconv.Atoi(os.Args[0])

	t1 := 0
	t2 := 1
	nextTerm := 0
	fmt.Print("Fibonacci Series :")
	for i := 1; i <= count; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Print(" ", nextTerm)
	}

}
