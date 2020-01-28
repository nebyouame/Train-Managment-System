package repository

import (
	"TrainProject/entity"
	"TrainProject/schedule"
	"github.com/jinzhu/gorm"
)

type BookGormRepo struct {
	conn *gorm.DB
}

func NewBookGormRepo(db *gorm.DB) schedule.BookRepository {
	return &BookGormRepo{conn:db}
}

func (bRepo *BookGormRepo) Books() ([]entity.Book, []error) {
	zss := []entity.Book{}
	errs := bRepo.conn.Find(&zss).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return zss, nil
}

func (bRepo *BookGormRepo) Book(id uint) (*entity.Book, []error) {
	zs := entity.Book{}
	errs := bRepo.conn.First(&zs, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &zs, nil
}

func (bRepo *BookGormRepo) UpdateBook(book *entity.Book) (*entity.Book, []error){
	bat := book
	errs := bRepo.conn.Save(bat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return bat, nil
}

func (bRepo *BookGormRepo) DeleteBook(id uint) (*entity.Book, []error) {
	bat, errs := bRepo.Book(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = bRepo.conn.Delete(bat, bat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return bat, nil
}

func (bRepo *BookGormRepo) StoreBook(book *entity.Book) (*entity.Book, []error) {
	bat := book
	errs := bRepo.conn.Create(bat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return bat, nil
}







































