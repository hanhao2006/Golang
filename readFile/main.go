package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// data,err :=ioutil.ReadFile("myfiletest.txt")
	// if err != nil{
	// 	fmt.Println("File reading error: ",err)
	// 	return
	// }
	// fmt.Println(string(data))

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

}
