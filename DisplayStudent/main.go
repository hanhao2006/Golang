// package main

// import "fmt"

// type Student struct {
// 	name  string
// 	grade int
// }

// func main() {

// 	var nb int

// 	fmt.Println("COLLEGE LASALLE")

// 	for ok := true; ok; ok = nb > 21 {
// 		fmt.Print("Enter the number of students (max:20): ")
// 		fmt.Scan(&nb)
// 	}
// 	stu := readAllStudent(nb)
// 	displayAllStu(stu)

// }

// func readOneStuName() string {
// 	var name string
// 	fmt.Print("name: ")
// 	fmt.Scan(&name)

// 	return name
// }

// func readOneStuGrade() int {
// 	var grade int
// 	for ok := true; ok; ok = grade < 0 || grade > 101 {
// 		fmt.Print("Grade: ")
// 		fmt.Scan(&grade)
// 	}
// 	return grade
// }

// func readAllStudent(nb int) []Student {
// 	stu := make([]Student, nb)
// 	for i := 0; i < nb; i++ {
// 		fmt.Println("Student ", i+1)
// 		stu[i].name = readOneStuName()
// 		stu[i].grade = readOneStuGrade()
// 	}
// 	return stu

// }

// func displayAllStu(stu []Student) {
// 	for _, s := range stu {
// 		fmt.Println(s.name, s.grade)
// 	}
// }
