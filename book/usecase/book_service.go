package usecase

import (
	"TrainSystem/book"
	"TrainSystem/entity"
)

type BookService struct {
	bookRepo book.BookRepostory
}

func NewBookService(bookRepository book.BookRepostory) book.BookService {
	return &BookService{bookRepo:bookRepository}
}

func (bs *BookService) Books() ([]entity.Book, []error) {
	bks, errs := bs.bookRepo.Books()
	if len(errs) > 0 {
		return nil, errs
	}
	return bks, errs
}

func (bs *BookService) Book(id uint) (*entity.Book, []error){
	bk, errs := bs.bookRepo.Book(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bs *BookService) CustomerBooks(customer *entity.User) ([]entity.Book, []error){
	bks, errs :=bs.bookRepo.CustomerBooks(customer)
	if len(errs) > 0 {
		return nil, errs
	}
	return bks,errs
}

func (bs *BookService) UpdateBook(book *entity.Book) (*entity.Book, []error){
	bk, errs := bs.bookRepo.UpdateBook(book)
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bs *BookService) DeleteBook(id uint) (*entity.Book, []error){
	bk, errs := bs.DeleteBook(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bs *BookService) StoreBook(book *entity.Book) (*entity.Book, []error){
	bk, errs := bs.StoreBook(book)
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}






























