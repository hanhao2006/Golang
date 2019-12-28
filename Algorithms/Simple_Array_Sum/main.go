package main

import "fmt"

func main() {

	arr := [6]int{1, 2, 3, 4, 10, 11}

	res := simpleArraySum(arr[:])
	fmt.Println(res)
}

func simpleArraySum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
