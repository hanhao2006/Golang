package main

import (
	"fmt"
)

func main() {
	var arr1 = [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)

	changarray(arr1)
	// array in Golang are value types
	fmt.Println(arr1)

	// when we want to change value, we need to use poiter.
	changarrayPoi(&arr1)

	fmt.Println(arr1)
}

func changarray(arr [5]int) {
	arr[0] = 125
}

func changarrayPoi(arr *[5]int) {
	arr[0] = 125
}
