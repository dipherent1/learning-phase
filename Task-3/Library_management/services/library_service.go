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
	book, book_not_exists := l.Books[book.Id]
	if book_not_exists {
		l.Books[book.Id] = book
		fmt.Printf("Book %v added\n\n", book.Title)
	}else{
		fmt.Printf("Book already exists")
	}
}

func (l *Library) AddMember(m models.Member) {
	l.Members[m.Id] = m
	fmt.Printf("member %v added\n", m.Id)
}

func (l *Library) RemoveBook(bookID int) {
	book, book_not_exist := l.Books[bookID]
	if book_not_exist {
		fmt.Printf("Book %v doesnt exist\n", book.Title)
		return
	}
	delete(l.Books, bookID)
	fmt.Printf("Book %v Removed\n", book.Title)

}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, book_not_exist := l.Books[bookID]

	if book_not_exist {
		return errors.New("book doesn't exist")

	} else if book.Status == "Borrowed" {
		return errors.New("book isn't available")

	}

	member, member_not_exists := l.Members[memberID]
	if member_not_exists {
		return errors.New("member doesn't exist")
	}

	book_address := &book
	temp := *book_address
	temp.Status = "Borrowed"
	// fmt.Print(temp.Status)

	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	fmt.Println("Book borrowed ")

	return nil

}

func (l *Library) ReturnBook(memberID int, bookID int) error {
	book, book_not_exist := l.Books[bookID]

	if book_not_exist {
		return errors.New("book doesn't exist")

	} else if book.Status != "Borrowed" {
		return errors.New("book not borrowed doesn't exist")

	}

	member, member_not_exists := l.Members[memberID]
	if member_not_exists {
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
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, member_not_exists := l.Members[memberID]
	if member_not_exists {
		return nil
	}
	return member.BorrowedBooks
}

func (l *Library) CreateMember(member models.Member) {
	member, member_not_exist := l.Members[member.Id]
	if member_not_exist {
		l.Members[member.Id] = member
		fmt.Printf("Member created\n\n")
	} else {
		fmt.Println("Member already exists")
	}
}
