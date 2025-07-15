package models


type status string

const (
	Borrowed  status = "borrowed"
	Available status = "available"
)

type Book struct{
	ID int
	Title string
	Author string
	Status status

}


