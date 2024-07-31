package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// check for palindrom
func is_palindrom(text string) {
	//declare left and right pointers
	l := 0
	r := len(text) - 1

	for {
		//check valid case
		if l > r {
			fmt.Println("Is palindrom")
			return
		}

		//check invalid case
		if text[l] != text[r] {
			fmt.Println("Is not palindrom")
			return
		}

		//move pointers
		l++
		r--
	}
}

// count word
func WordCount(s string) map[string]int {
	count := make(map[string]int)
	for _, word := range strings.Fields(s) {
		_, ok := count[word]
		if ok {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}
	return count
}

var choice string

func main() {
	fmt.Print("Enter 1 or 2\n 1. Palindrom\n 2. Word count\n ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	// get choice string

	for {
		choice, err := reader.ReadString('\n')
		//remove new line
		choice = strings.TrimSpace(choice)
		//check for validity
		if err != nil || (choice != "1" && choice != "2") {
			fmt.Println("Please enter 1 or 2")

		} else {
			//get text
			txt, _ := reader.ReadString('\n')
			txt = strings.TrimSpace(txt)

			txt = ConvertForPalindrom(txt)
			if choice == "1" {
				is_palindrom(txt)
			} else {
				fmt.Print(WordCount(txt))
			}

			break
		}
	}
}

func ConvertForPalindrom(txt string) string {
	var temp string

	if choice == "1" {
		for _, l := range txt {
			if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') || (l >= '0' && l <= '9') {
				temp += string((l))
			}
		}
	} else {
		for _, l := range txt {
			if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') || (l >= '0' && l <= '9') || (l == ' ') {
				temp += string((l))
			}
		}

	}

	return temp

}
