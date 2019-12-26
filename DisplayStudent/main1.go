package main

import "fmt"

type Person struct {
	name string
	age  int
}

func printer(arr []int) {
	arr[1] = 10

	fmt.Println(arr)
}

func main() {

	arr := []int{1, 2, 3}

	fmt.Println(arr)

	printer(arr)

	fmt.Println(arr)

}
