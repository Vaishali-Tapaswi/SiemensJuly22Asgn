package mainimport "fmt"func main() {
	var x int
	var y int
	fmt.Println("Enter two numbers")
	fmt.Scanln(&x, &y)s, c := function1(x)
	fmt.Println("Square and cube of x", s, c)
	s1, c1 := function1(y)
	fmt.Println("Square and cube of y", s1, c1)
}func function1(x int) (int, int) {
	var a = x * x
	return a, a * x
}