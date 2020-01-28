package repository

import (
	"TrainProject/entity"
	"TrainProject/schedule"
	"github.com/jinzhu/gorm"
)

type RoleGormRepo struct {
	conn *gorm.DB
}

func NewRoleGormRepo(db *gorm.DB) schedule.RoleRepository {
	return &RoleGormRepo{conn:db}
}

func (rRepo *RoleGormRepo) Roles() ([]entity.Role, []error) {
	schs := []entity.Role{}
	errs := rRepo.conn.Find(&schs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return schs, nil
}

func (rRepo *RoleGormRepo) Role(id uint) (*entity.Role, []error) {
	sch := entity.Role{}
	errs := rRepo.conn.First(&sch, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return &sch, nil
}

func (rRepo *RoleGormRepo) UpdateRole(role *entity.Role) (*entity.Role, []error){
	sat := role
	errs := rRepo.conn.Save(sat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (rRepo *RoleGormRepo) DeleteRole(id uint) (*entity.Role, []error) {
	sat, errs := rRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = rRepo.conn.Delete(sat, sat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (rRepo *RoleGormRepo) StoreRole(role *entity.Role) (*entity.Role, []error) {
	sat := role
	errs := rRepo.conn.Create(sat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}



