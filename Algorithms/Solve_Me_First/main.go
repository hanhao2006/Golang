package main

import "fmt"

func main() {

	x := 0
	y := 0

	fmt.Print("Enter your filst number")
	fmt.Scan(&x)
	fmt.Print("Enter your second number")
	fmt.Scan(&y)

	sum := sum(x, y)

	fmt.Println(sum)

}

func sum(x int, y int) int {
	return x + y
}
