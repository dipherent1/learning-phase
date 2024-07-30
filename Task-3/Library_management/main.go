package main

import (
	"fmt"
	"libmgr/controllers"
)

func main() {
	for {
		fmt.Printf("\n\n\n---Library Management System---\n\n")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Create member")
		fmt.Println("8. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			controllers.AddBook()
		case 2:
			controllers.RemoveBook()
		case 3:
			controllers.BorrowBook()
		case 4:
			controllers.ReturnBook()
		case 5:
			controllers.ListAvailableBooks()
		case 6:
			controllers.ListBorrowedBooks()
		case 7:
			controllers.CreateMember()
		case 8:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
