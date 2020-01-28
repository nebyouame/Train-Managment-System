package service

import (
	"TrainProject/entity"
	"TrainProject/schedule"
)

type Bookservice struct {
	bookRepo schedule.BookRepository
}

func NewBookService(bookRepository schedule.BookRepository) schedule.BookService{
	return &Bookservice{bookRepo:bookRepository}
}

func (bs *Bookservice) Books() ([]entity.Book, []error) {
	books, errs := bs.bookRepo.Books()
	if len(errs) > 0 {
		return nil, errs
	}

	return books, nil
}

func (bs *Bookservice) StoreBook(book *entity.Book) (*entity.Book, []error){
	bk, errs := bs.bookRepo.StoreBook(book)
	if len(errs) > 0 {
		return nil, errs
	}

	return bk, nil
}

func (bs *Bookservice) Book(id uint) (*entity.Book, []error) {
	b, errs := bs.bookRepo.Book(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return b, nil
}

func (bs *Bookservice) UpdateBook(book *entity.Book) (*entity.Book, []error){
	bak, errs := bs.bookRepo.UpdateBook(book)
	if len(errs) > 0 {
		return nil, errs
	}

	return bak, nil
}

func (bs *Bookservice) DeleteBook(id uint) (*entity.Book, []error)  {
	bak, errs := bs.bookRepo.DeleteBook(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return bak, nil
}

















