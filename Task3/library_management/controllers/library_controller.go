package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
	"strings"
)

func RegisterBook(l *services.Library) {
	var title, author string


	fmt.Println("Enter the book title: ")
	fmt.Scanln(&title)

	fmt.Println("Enter the book author: ")
	fmt.Scanln(&author)


	book := models.Book{
		Title:  title,
		Author: author,
		Status: "available",
	}

	l.AddBook(book)
	fmt.Println("Book registered successfully!")
}

func RegisterMember(l *services.Library) {
	var name string

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	

	member := models.Member{
		Name:          name,
		BorrowedBooks: []models.Book{},
	}
	l.AddMember(member)

}

func RemoveBook(l *services.Library) {
	var id int
	fmt.Println("Enter the book ID you want to remove: ")
	fmt.Scanln(&id)
	l.RemoveBook(id)
}

func BorrowBook(l *services.Library) {
	var bookID, id int
	fmt.Println("Enter book ID to borrow: ")
	fmt.Scanln(&bookID)
	fmt.Println("Enter your member ID: ")
	fmt.Scanln(&id)
	err := l.BorrowBook(bookID, id)
	if err != nil {
		fmt.Println(err)
	} 
}

func ReturnBook(l *services.Library) {
	var bookID, id int
	fmt.Println("Enter book ID to return: ")
	fmt.Scanln(&bookID)
	fmt.Println("Enter your member ID: ")
	fmt.Scanln(&id)
	err := l.ReturnBook(bookID, id)
	if err != nil {
		fmt.Println( err)
	} 
}

func ListAvailableBooks(l *services.Library) {
	books := l.ListAvailableBooks()
	if books == nil {
		fmt.Println("There are no available books to list")
		return
	}
	fmt.Println("\nAvailable Books:")
	fmt.Printf("ID \t | Title \t | Author\t\n")
	fmt.Println(strings.Repeat("-", 60))

	for _, book := range books {
		fmt.Printf("%v\t | %v\t | %v\t\n", book.ID, book.Title, book.Author)
	}
	
}

func ListBorrowedBooks(l *services.Library) {
	id := 0
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&id)

	borrowed := l.ListBorrowedBooks(id)
	member := l.Members[id]

	if len(borrowed) == 0 {
		fmt.Println("This member has not borrowed any books.")
		return
	}

	fmt.Printf("\nBooks borrowed by %s:\n", member.Name)
	fmt.Printf("ID \t | Title \t | Author\t\n")
	fmt.Println(strings.Repeat("-", 60))

	for _, book := range borrowed{
		fmt.Printf("%v\t | %v\t | %v\t\n", book.ID, book.Title, book.Author)
	}
}

