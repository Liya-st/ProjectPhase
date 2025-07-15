package main

import (
	"fmt"
	"library_management/controllers"
	"library_management/services"
	"library_management/models"
)

func main() {
	library := &services.Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}

	for {
		fmt.Println("\n Library Management System")
		fmt.Println("1. Register a Book")
		fmt.Println("2. Register a Member")
		fmt.Println("3. List Available Books")
		fmt.Println("4. List Borrowed Books by Member")
		fmt.Println("5. Borrow a Book")
		fmt.Println("6. Return a Book")
		fmt.Println("7. Remove a Book")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			controllers.RegisterBook(library)
		case 2:
			controllers.RegisterMember(library)
		case 3:
			controllers.ListAvailableBooks(library)
		case 4:
			controllers.ListBorrowedBooks(library)
		case 5:
			controllers.BorrowBook(library)
		case 6:
			controllers.ReturnBook(library)
		case 7:
			controllers.RemoveBook(library)
		case 8:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
