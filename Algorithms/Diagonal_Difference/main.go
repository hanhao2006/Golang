package main

import "fmt"

import "math"

// 11 2 4
// 4  5 6
// 10 8 -12
// row[0]->[0][1][2] row[1]->[0][1][2] row[2]->[0][1][2]
// row[0][0] + row[1][1] + row[2][2]
// row[0][2] + row[1][1] + row[2][0]

func main() {
	arr := [3][3]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	sum := addlefttoright(arr)
	fmt.Println(sum)
	sum1 := addrighttoleft(arr)
	fmt.Println(sum1)
	absSumandSum1 := abs(sum, sum1)
	fmt.Println(absSumandSum1)

	//==============================================================
	sumFromLeft, sumFromRight := addLefttoRightANDop(arr)
	fmt.Println(sumFromLeft, sumFromRight)

}

func addLefttoRightANDop(ar [3][3]int) (int, int) {
	sumFromLeft := 0
	sumFromRight := 0
	for i := 0; i < len(ar); i++ {
		sumFromLeft += ar[i][i]
		sumFromRight += ar[len(ar)-1-i][i]
	}
	return sumFromLeft, sumFromRight
}

//=====================================================
func addlefttoright(ar [3][3]int) int {
	sum := 0
	for i := 0; i < len(ar); i++ {
		sum += ar[i][i]
	}
	return sum
}

func addrighttoleft(ar [3][3]int) int {
	sum := 0
	for i := 2; i >= 0; i-- {
		sum += ar[len(ar)-1-i][i]
	}
	return sum
}

func abs(a int, b int) float64 {

	return math.Abs(float64(a) - float64(b))
}
