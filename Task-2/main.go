package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//check for palindrom
func is_palindrom(text string) {
	//declare left and right pointers
	l :=0
	r := len(text)

	for{
		//check invalid case
		if text[l] != text[r]{
			fmt.Println("Is not palindrom")
			return
		}
		
		//check valid case
		if l>r{
			fmt.Println("Is palindrom")
			return 
		}
		
		//move pointers
		l++
		r--
	}
}

//count word
func wordCount(text string) {
	// count = make(map[string] int)


}

func main() {
	fmt.Print("Enter 1 or 2\n 1. Palindrom\n 2. Word count\n ")
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter number: ")
	var text string
	// get text string
	for{
		text, err := reader.ReadString('\n')
		//remove new line
		text = strings.TrimSpace(text)
		//check for validity
		if err != nil || (text != "1" && text!="2") {
			fmt.Println("Please enter 1 or 2")
			
		}else{
			break
		}
	}
	
	if text == "1"{
		is_palindrom(text)
	} else{
		wordCount(text)
	}

	
}