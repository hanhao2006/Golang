package main

import (
	"fmt"
	"strings"
)

func main() {

	options := []string{"Addition", "Division", "Substraction", "Mulitiplication", "Quit"}
	re := " "
	
	for i := true; i; i = strings.ToLower(re) == "n"{	
		displayHeader("CACLATOR", "By Hao Han")
		displayMenu(options)
		selectoption := selectoptions(len(options))
		
		for i := true; i; i = strings.ToLower(re) == "y" {

			nb := displayCalculatormanu(options[selectoption-1])

			calc := Cal{readvalues(nb), 0}
			switch selectoption {
			case 1:
				calc.Addition()
				fmt.Println("Answer: ", calc.result)
			case 2:
				calc.Division()
				fmt.Printf("Answer: %.2f\n", calc.result)
			case 3:
				calc.Substraction()
				fmt.Println("Answer: ", calc.result)
			case 4:
				calc.Mulitiplication()
				fmt.Println("Answer: ", calc.result)
			case 5:
				fmt.Println("Quit")

			}
			restart(&re)
		}
	}
}
