package service

import (
	"TrainProject/entity"
	"TrainProject/schedule"
)

type RoleService struct {
	roleRepo schedule.RoleRepository
}

func NewRoleService(roleRepository schedule.RoleRepository) schedule.RoleService {
	return &RoleService{roleRepo:roleRepository}
}

func (is *RoleService) Roles() ([]entity.Role, []error){
	infs, errs := is.roleRepo.Roles()
	if len(errs) > 0 {
		return nil, errs
	}

	return infs, nil

}

func (is *RoleService) Role(id uint) (*entity.Role, []error){
	infs, errs := is.roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return infs, nil

}

func (is *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	inf, errs := is.roleRepo.UpdateRole(role)
	if len(errs) > 0 {
		return nil, errs
	}

	return inf, nil
}

func (is *RoleService) DeleteRole(id uint) (*entity.Role, []error) {
	inf, errs := is.roleRepo.DeleteRole(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return inf, nil
}

func (is *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {
	inf, errs := is.roleRepo.StoreRole(role)
	if len(errs) > 0 {
		return nil, errs
	}

	return inf, nil
}





















