package main

import "fmt"

//declare map type to store user profile
type subject_grades map[string] float64

//declare sum method
func (profile subject_grades) avg() float64{
	//declare variables
	total := 0.0
	count := 0.0
	
	for _,val:=range profile{
		
		total+=val
		count++
		
	}
	return total/count

}

//declare show method
func (profile subject_grades) show() {
	for key,val := range profile{
		fmt.Printf("%v : %v \n",key,val)
	}
	fmt.Printf("Average : %v",profile.avg())

}


func main() {
	//declare variables
	var name string
	var subjectNum int
	var profile = make(subject_grades)
	
	//ask user for name
	fmt.Println("Please write your name!")
	fmt.Scanln(&name)
	
	
	//ask user for number of subj
	fmt.Printf("Hi %v ",name + " " + "Please write the number of subjects you take! ")
	fmt.Scanln(&subjectNum)

	//declare variables
	var subject string
	var grade float64
	

	//ask user of subject name and grade
	for i := 0; i < subjectNum;{
		
		fmt.Println("Please enter your subject name")
		fmt.Scan(&subject)
		fmt.Println("Please enter your grade")
		fmt.Scan(&grade)
		
		//check for validity
		if grade < 0 || grade> 100{
			fmt.Println("Grade must be from 0 to 100 please try again")
		}else {

			profile[subject] = grade
			i++
		}
	
	}
	fmt.Printf("Usename : %v \n",name)
	profile.show()

}
