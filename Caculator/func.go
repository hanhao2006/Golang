package main

import (
	"fmt"
	"strings"
)

func displayHeader(text string, subtitile string) {
	fmt.Println(text)
	fmt.Println(strings.Repeat("_", len(text)))
	fmt.Println(subtitile)
	fmt.Println(strings.Repeat("*", len(subtitile)))
}

func displayMenu(options []string) {
	for i, option := range options {
		fmt.Printf("%d - %s \n", i+1, option)
	}
}

func selectoptions(lenoptions int) int {
	var optionselect int
	for i := true; i; i = optionselect < 0 || optionselect > lenoptions {
		fmt.Printf("Enter your select: 1 - %d: ", lenoptions)
		if _, err := fmt.Scan(&optionselect); err != nil {
			fmt.Println(err)
			optionselect = selectoptions(lenoptions)
		}
	}
	return optionselect
}

func displayCalculatormanu(option string) int {
	displayHeader("Calculator", option)
	var v int
	fmt.Print("Enter the number of values of operate: ")
	if _, err := fmt.Scan(&v); err != nil {
		fmt.Println(err)
	}
	return v
}

func readvalues(v int) []float64 {
	values := make([]float64, v)
	for i := 0; i < v; i++ {
		fmt.Print("Number: ", i+1, "  ")
		if _, err := fmt.Scan(&values[i]); err != nil {
			fmt.Println(err)
		}
	}
	return values
}

// type Calculator interface {
// 	Addition()
// }

type Cal struct {
	//num1, num2 float64
	values []float64
	result float64
}

func (a *Cal) Addition() float64 {
	for _, val := range a.values {
		a.result += val
	}
	return a.result
}

func (a *Cal) Division() float64 {
	a.result = a.values[0]
	for i := 1; i < len(a.values); i++ {
		a.result /= a.values[i]
	}
	return a.result
}

func (a *Cal) Substraction() float64 {
	a.result = a.values[0]
	for i := 1; i < len(a.values); i++ {
		a.result -= a.values[i]
	}
	return a.result
}

func (a *Cal) Mulitiplication() float64 {
	a.result = 1
	for i := 0; i < len(a.values); i++ {
		a.result -= a.values[i]
	}
	return a.result
}

func restart(re *string){
	tempRe :=""
	fmt.Print("Restart ? : ")
	if _, err := fmt.Scan(&tempRe); err !=nil {
		fmt.Println(err)
	}
	*re = tempRe
}
