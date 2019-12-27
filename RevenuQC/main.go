package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	name := readUserName()
	gender := readgender()
	var manyChild int
	check := true
	// fmt.Println(name)
	// fmt.Println(gender)
	if gender == "F" {
		fmt.Println("Mandan ", name)
	} else {

		fmt.Println("Sir ", name)
	}

	for ok := true; ok; ok = check == false && manyChild > 6 {
		fmt.Print("How many child: ")
		fmt.Scan(&manyChild)
		check = IsInt(string(manyChild))
	}
	fmt.Println(manyChild)
}

// IsInt Check is integer
func IsInt(s string) bool {
	l := len(s)

	if strings.HasPrefix(s, "-") {
		l = l - 1
		s = s[1:]
	}

	reg := fmt.Sprintf("\\d{%d}", l)

	rs, err := regexp.MatchString(reg, s)

	if err != nil {
		return false
	}

	return rs
}

func readUserName() string {
	name := bufio.NewScanner(os.Stdout)
	check := false
	for ok := true; ok; ok = check == true {

		fmt.Print("Enter your fullname: ")
		name.Scan()
		check = IsInt(name.Text())
	}
	return name.Text()
}

func readgender() string {
	var gender string
	var upper string

	for ok := true; ok; ok = gender != "f" && gender != "F" && gender != "m" && gender != "M" {
		fmt.Print("Select gender (f/m): ")
		fmt.Scan(&gender)
		upper = strings.ToUpper(gender)

	}

	return upper
}

// func main() {

// 	str1 := "1234"

// 	/** converting the str1 variable into an int using Atoi method */
// 	i1, err := strconv.Atoi(str1)
// 	if err == nil {
// 		fmt.Println(i1)
// 	}

// 	str2 := "5678"
// 	/** converting the str2 variable into an int using ParseInt method */
// 	i2, err := strconv.ParseInt(str2, 10, 64)
// 	if err == nil {
// 		fmt.Println(i2)
// 	}
// }
