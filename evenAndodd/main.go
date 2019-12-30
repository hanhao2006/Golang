package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, k := range arr {

		if k%2 == 0 {
			fmt.Println(k, " is even")
		} else {
			fmt.Println(k, " is odd")
		}

	}

	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, k := range arr1 {
		if k%2 == 0 {
			fmt.Printf("%v is even\n", k)
			continue
		}
		fmt.Printf("%v is odd\n", k)
	}
}
