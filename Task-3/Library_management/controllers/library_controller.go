package controllers

import (
	"fmt"
	"libmgr/models"
	"libmgr/services"
)

var library = services.NewLibrary()

func AddBook() {
	var id int
	var title, author string

	fmt.Print("Enter book ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter book title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter book author: ")
	fmt.Scanln(&author)

	library.AddBook(models.Book{Id: id, Title: title, Author: author, Status: "Available"})
}

func RemoveBook() {
	var id int
	fmt.Print("Enter book ID to remove: ")
	fmt.Scanln(&id)

	if e:=library.RemoveBook(id); e != nil{
		fmt.Println("Error",e)
	}

}

func BorrowBook() {
	var bookID int
	fmt.Print("Enter book ID to borrow: ")
	fmt.Scanln(&bookID)

	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	if e :=library.BorrowBook(bookID, memberID); e != nil{
		fmt.Println("Error",e)
	}
	
}

func ReturnBook() {
	var bookID int
	fmt.Print("Enter book ID to return: ")
	fmt.Scanln(&bookID)
	
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)
	
	
	if e := library.ReturnBook(bookID, memberID); e != nil{
		fmt.Println("Error",e)
	}

}

func ListAvailableBooks() {
	books := library.ListAvailableBooks()
	fmt.Printf("Available Books:\n\n")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)

	}

}

func ListBorrowedBooks() {
	var memberID int

	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	books := library.ListBorrowedBooks(memberID)
	fmt.Printf("Borrowed Books by Member %d:\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	}
}

func CreateMember(){
	var Id int
	fmt.Println("Enter ID")
	fmt.Scan(&Id)
	
	var Name string
	fmt.Println("Enter Name")
	fmt.Scan(&Name)
	
	var BorrowedBooks []models.Book 

	if e:=library.CreateMember(models.Member{Id: Id, Name: Name, BorrowedBooks: BorrowedBooks}); e!=nil{
		fmt.Println("Error",e)
	}
	

}
