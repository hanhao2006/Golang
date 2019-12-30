package main

import "fmt"

//Given an array of integers, calculate the fractions of its elements that are positive, negative, and are zeros.
//Print the decimal value of each fraction on a new line.
// -4,-9 =>2
// 3,4, 1, => 3
// 0 = 1
// total 6
func main() {
	arr := [...]int{-4, 3, -9, 0, 4, 1}

	negative, positive, zero := findNumber(arr)
	nres, pres, zres := cal(len(arr), negative, positive, zero)

	fmt.Printf("%d\n", len(arr))
	fmt.Printf("%f\n", nres)
	fmt.Printf("%f\n", pres)
	fmt.Printf("%f\n", zres)

}

func findNumber(arr [6]int) (int, int, int) {
	positive := 0
	negative := 0
	zero := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative++
		} else if arr[i] > 0 {
			positive++
		} else {
			zero++
		}

	}
	return negative, positive, zero
}

func cal(lenght int, negative int, positive int, zero int) (float64, float64, float64) {
	var (
		nres float64
		pres float64
		zres float64
	)

	nres = float64(negative) / float64(lenght)
	pres = float64(positive) / float64(lenght)
	zres = float64(zero) / float64(lenght)

	return nres, pres, zres

}
