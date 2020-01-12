package book

import "TrainSystem/entity"

type BookRepostory interface {
	Books() ([]entity.Book, []error)
	Book(id uint) (*entity.Book, []error)
	CustomerBooks(customer *entity.User) ([]entity.Book, []error)
	UpdateBook(book *entity.Book) (*entity.Book, []error)
	DeleteBook(id uint) (*entity.Book, []error)
	StoreBook(book *entity.Book) (*entity.Book, []error)
}