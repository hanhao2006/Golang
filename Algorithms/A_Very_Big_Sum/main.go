package main

import "fmt"

import "bufio"

import "os"

func main() {

	var nb int
	fmt.Println("Enter number of the values: ")
	fmt.Scan(&nb)
	readValueList()

}

func readValueList() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}
