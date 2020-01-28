package repository

import (
	"TrainProject/entity"
	"TrainProject/schedule"
	"github.com/jinzhu/gorm"
)

type UserGormRepo struct {
	conn *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) schedule.UserRepository {
	return &UserGormRepo{conn:db}
}

func (uRepo *UserGormRepo) Users() ([]entity.User, []error) {
	cks := []entity.User{}
	errs := uRepo.conn.Find(&cks).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return cks, nil
}

func (uRepo *UserGormRepo) User(id uint) (*entity.User, []error) {
	ck := entity.User{}
	errs := uRepo.conn.First(&ck, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return &ck, nil
}

func (uRepo *UserGormRepo) UpdateUser(user *entity.User) (*entity.User, []error){
	ck := user
	errs := uRepo.conn.Save(ck).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return ck, nil
}

func (uRepo *UserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
	ck, errs := uRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = uRepo.conn.Delete(ck, ck.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return ck, nil
}

func (uRepo *UserGormRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	ck := user
	errs := uRepo.conn.Create(ck).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return ck, nil
}


