# Library Management System Documentation

## Overview
This system is a console-based library management application written in Go. It allows users to manage books and members, including adding, removing, borrowing, and returning books.

## Features
- **Add a New Book**: Add new books to the library's collection.
- **Remove an Existing Book**: Remove books from the library by ID.
- **Borrow a Book**: Allows members to borrow available books.
- **Return a Book**: Allows members to return borrowed books.
- **List Available Books**: Displays all books that are currently available for borrowing.
- **List Borrowed Books**: Shows all books borrowed by a specific member.
- **Add member**: Add member.

## Structs
- **Book**
  - `ID` (int)
  - `Title` (string)
  - `Author` (string)
  - `Status` (string)

- **Member**
  - `ID` (int)
  - `Name` (string)
  - `BorrowedBooks` ([]Book)

## Interface
- **LibraryManager**
  - `AddBook(book Book)`
  - `RemoveBook(bookID int)`
  - `BorrowBook(bookID int, memberID int) error`
  - `ReturnBook(bookID int, memberID int) error`
  - `ListAvailableBooks() []Book`
  - `ListBorrowedBooks(memberID int) []Book`
  - `CreateMember(member Member)`

## Implementation
- **Library**: Implements the `LibraryManager` interface.
- **LibraryManager Interface**: Defined in `services/library_service.go`.
- **Console Controller**: Defined in `controllers/library_controller.go`.

## Error Handling
The system handles errors such as:
- Attempting to borrow an unavailable or non-existent book.
- Attempting to return a book not borrowed by the member.
- Attempting to remove a non-existent book.

## Usage
1. **Run the Application**: Start the application by running `go run main.go`.
2. **Follow the Menu**: Use the menu to navigate through the functionalities.


## Dependencies
- This project requires Go to be installed. No external dependencies are needed beyond the standard library.

## License
This project is open-source and available under the MIT License.

