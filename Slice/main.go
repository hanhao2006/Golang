package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// A Slice is just a reference to an underlying array

	slice := arr[2:]
	fmt.Println(slice)
	fmt.Println(arr)

	changeslice(slice)
	fmt.Println(slice)
	fmt.Println(arr)
	fmt.Println("=========other slice====================")

	slice2 := arr[4:]
	fmt.Println(slice2)
	// The capacity is the number of elements in the underlying array starting from the first element in the slice.
	//You can find the length and capacity of a slice using the built-in functions len() and cap()
	fmt.Println(arr)
	slice3 := arr[2:4]
	fmt.Println(slice3)
	slice4 := slice3[2:4]
	fmt.Println(slice4)

}

func changeslice(slice []int) {
	slice[2] = 125
}
