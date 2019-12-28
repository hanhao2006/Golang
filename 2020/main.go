package main

import "fmt"

func main() {
	year2019 := map[string]int{"c++": 4, "html": 4, "css": 5, "sql": 8, "c#": 8, "js": 9, "php": 7, "ios": 4, "java": 7, "python": 1}
	learnagain := ListLearnMore(year2019)
	fmt.Println("\n========2019 studied========")
	dispaplyDic(year2019)
	year2020List := [...]string{"Go", "React", "internship"}
	fmt.Println("===========2020 List==================")
	displayList(year2020List[:], "2020 List")
	newyear2020List := append(year2020List[:])
	for _, li := range learnagain {
		newyear2020List = append(newyear2020List, li)
	}
	fmt.Println("=========2020 have to DO List==========")
	displayList(newyear2020List, "TO DO LIST")
	fmt.Println("=======HAPPY NEW YEAR :( ==============")
	changlisttomap(learnagain, year2019, year2020List[:])
}
//ListLearnMore got used to uppercase in the first letter
func ListLearnMore(list map[string]int) []string {
	learnlearn := []string{}
	for k, v := range list {
		//fmt.Println(k, v)
		if v > 4 {
			learnlearn = append(learnlearn, k)
		}
	}
	return learnlearn
}
func changlisttomap(list []string, dicarr map[string]int, newlist []string) {
	dic := map[string]int{}
	for _, l := range list {
		for k, v := range dicarr {
			if l == k {
				dic[l] = v
			}
		}
	}
	for _, l := range newlist {
		dic[l] = 0
	}
	dispaplyDic(dic)
}
func displayList(list []string, text string) {
	for _, l := range list {
		fmt.Println(text, " ", l)
	}
}
func dispaplyDic(list map[string]int) {
	for k, v := range list {
		fmt.Printf("langang %v Passion %d%%\n", k, v)
	}
}
