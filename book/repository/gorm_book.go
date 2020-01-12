package repository

import (
	"TrainSystem/book"
	"TrainSystem/entity"
	"github.com/jinzhu/gorm"
)

type BookGormRepo struct {
	conn *gorm.DB
}

func NewBookGormRepo(db *gorm.DB) book.BookRepostory {
	return &BookGormRepo{conn:db}
}
func(bookRepo *BookGormRepo) Books() ([]entity.Book, []error){
	books := []entity.Book{}
	errs := bookRepo.conn.Find(&books).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return books, errs
}
func(bookRepo *BookGormRepo) Book(id uint) (*entity.Book, []error){
	book := entity.Book{}
	errs := bookRepo.conn.First(&book, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &book, errs
}

func (bookRepo *BookGormRepo) UpdateBook(book *entity.Book) (*entity.Book, []error){
	bk := book
	errs := bookRepo.conn.Save(bk).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bookRepo *BookGormRepo) DeleteBook(id uint)(*entity.Book, []error){
	bk, errs := bookRepo.Book(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = bookRepo.conn.Delete(bk, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bookRepo *BookGormRepo) StoreBook(book *entity.Book) (*entity.Book, []error){
	bk := book
	errs := bookRepo.conn.Create(bk).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bookRepo *BookGormRepo) CustomerBooks(customer *entity.User) ([]entity.Book, []error){
	custBooks := []entity.Book{}
	errs := bookRepo.conn.Model(customer).Related(&custBooks, "Books").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return custBooks, errs
}













































