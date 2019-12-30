package main

import "fmt"

func main() {
	a := 2
	b := 1

	a = a ^ b
	fmt.Println(a)
	b = a ^ b
	fmt.Println(a, b)
	a = a ^ b
	fmt.Println(a, b)

	// if a = a > b => a + b   if a = a < b => b - a

}
