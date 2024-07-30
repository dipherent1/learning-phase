package services

import (
	"errors"
	"fmt"
	"libmgr/models"
)

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member)}
}

func (l *Library) AddBook(book models.Book) {
	_, exist := l.Books[book.Id]
	if !exist {
		l.Books[book.Id] = book
		fmt.Printf("Book %v added\n\n", book.Title)
	}else{
		fmt.Printf("Book already exists")
	}
}


func (l *Library) RemoveBook(bookID int) error{
	book, exist := l.Books[bookID]
	if !exist {
		return errors.New("book doesn't exist")
		
	}
	delete(l.Books, bookID)
	fmt.Printf("Book %v Removed\n", book.Title)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exist := l.Books[bookID]

	if !exist {
		return errors.New("book doesn't exist")

	} else if book.Status == "Borrowed" {
		return errors.New("book isn't available")

	}

	member, exist := l.Members[memberID]
	if !exist {
		return errors.New("member doesn't exist")
	}


	// fmt.Print(temp.Status)
	
	book.Status = "Borrowed"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	fmt.Println("Book borrowed ")

	return nil

}

func (l *Library) ReturnBook(memberID int, bookID int) error {
	book, exist := l.Books[bookID]

	if !exist {
		return errors.New("book doesn't exist")

	} else if book.Status != "Borrowed" {
		return errors.New("book not borrowed doesn't exist")

	}

	member, exist := l.Members[memberID]
	if !exist {
		return errors.New("member doesn't exist")

	}

	for i, b := range member.BorrowedBooks {
		if b.Id == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	book.Status = "Available"
	l.Books[bookID] = book
	l.Members[memberID] = member
	fmt.Println("book returned")
	return nil

}

func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	// fmt.Print(l.Books)
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exist := l.Members[memberID]
	if !exist {
		return nil
	}
	return member.BorrowedBooks
}

func (l *Library) CreateMember(member models.Member) error {
	_, exist := l.Members[member.Id]
	if !exist {
		l.Members[member.Id] = member
		fmt.Printf("Member created\n\n")
	} else {
		return errors.New("member already exist")
	}
	return nil
}
