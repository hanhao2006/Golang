package main

import "net/http"

import "fmt"

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checklink(link, c)
		fmt.Println(<-c)

	}

}

func checklink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down"
		return
	}
	fmt.Println(link, "is up")
	c <- "Yep is up"
}
