package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	AddMember(member models.Member)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	incrementBookID() int
	incrementMemberID()int

}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
	BookID int
	MemberID int
}

func (l *Library) incrementBookID() int {
	l.BookID++
	return l.BookID
}

func (l *Library) incrementMemberID() int {
	l.MemberID++
	return l.MemberID
}

func (l *Library) AddBook(book models.Book) {
	id := l.incrementBookID()
	book.ID = id
	l.Books[book.ID] = book
	fmt.Printf("Book added: %s (ID: %d)\n", book.Title, book.ID)
}

func (l *Library) AddMember(member models.Member) {
	id := l.incrementMemberID()
	member.ID = id
	l.Members[id] = member
	fmt.Printf("Member added: %s (ID: %d)\n", member.Name, member.ID)
}

func (l *Library) RemoveBook(bookID int) {
	if _, ok := l.Books[bookID]; !ok {
		fmt.Printf("There is no book with that ID. ID: %d\n", bookID)
		return
	}
	delete(l.Books, bookID)
	fmt.Printf("Book with ID %d removed successfully.\n", bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book not found: id %d", bookID)
	}

	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member not found: id %d", memberID)
	}

	if book.Status == models.Borrowed {
		return errors.New("book is already borrowed")
	}

	book.Status = models.Borrowed
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	fmt.Printf("Book with ID %d borrowed by member %s (ID: %d)\n", bookID, member.Name, memberID)

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book not found: id %d", bookID)
	}

	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member not found: id %d", memberID)
	}

	if book.Status == models.Available {
		return errors.New("book was not borrowed")
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	book.Status = models.Available
	l.Books[bookID] = book
	l.Members[memberID] = member
	fmt.Printf("Book with ID %d returned by member %s (ID: %d)\n", bookID, member.Name, memberID)

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.Books {
		if book.Status == models.Available {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.Members[memberID]
	if !ok {
		fmt.Printf("There is no member with that ID. ID: %d\n", memberID)
		return nil
	}
	return member.BorrowedBooks
}
