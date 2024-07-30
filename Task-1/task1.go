package main

import "fmt"

// var subject_grade map[string] int

func main() {
	//declare variables
	var name string
	var subjectNum int

	fmt.Println("Please write your name!")
	fmt.Scanln(&name)

	fmt.Println("Please write the number of subjects you take!")
	fmt.Scanln(&subjectNum)

	var subject string
	var grade int
	for i := 0; i < subjectNum; i++ {
		fmt.Println("Please enter your subject name and grade")
		fmt.Scan(&subject, &grade)

	}
}
