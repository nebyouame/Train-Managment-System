package service

import (
	"TrainProject/entity"
	"TrainProject/schedule"
)

type UserService struct {
	userRepo schedule.UserRepository
}

func NewUserService(userRepository schedule.UserRepository) schedule.UserService {
	return &UserService{userRepo:userRepository}
}

func (us *UserService) Users() ([]entity.User, []error){
	uses, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return uses, nil
}

func (us *UserService) User(id uint) (*entity.User, []error){
	uses, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return uses, nil
}

func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error){
	use, errs := us.userRepo.UpdateUser(user)
	if len(errs) > 0 {
		return nil, errs
	}

	return use, nil
}

func (us *UserService) DeleteUser(id uint) (*entity.User, []error){
	use, errs := us.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return use, nil
}


func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	use, errs := us.userRepo.StoreUser(user)
	if len(errs) > 0 {
		return nil, errs
	}

	return use, nil
}


















































