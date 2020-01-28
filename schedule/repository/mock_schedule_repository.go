package repository

import (
	"TrainProject/entity"
	"TrainProject/schedule"
	"errors"
	"github.com/jinzhu/gorm"
)

type MockScheduleRepo struct {
	conn *gorm.DB
}

func  NewMockScheduleRepo(db *gorm.DB) schedule.ScheduleRepository {
	return &MockScheduleRepo{conn:db}
}

func(mSchRepo *MockScheduleRepo) Schedules() ([]entity.Schedule, []error) {
	schs := []entity.Schedule{entity.ScheduleMock}
	return schs, nil
}

func (mSchRepo *MockScheduleRepo) Schedule(id uint) (*entity.Schedule, []error) {
	sch := entity.ScheduleMock
	if id==1 {
		return &sch, nil
	}
	return nil, []error{errors.New("Not found")}
}

func (mSchRepo *MockScheduleRepo) UpdateSchedule(schedule *entity.Schedule) (*entity.Schedule, []error) {
	sat := entity.ScheduleMock
	return &sat, nil
}

func(mSat *MockScheduleRepo) DeleteSchedule(id uint) (*entity.Schedule, []error) {
	sat := entity.ScheduleMock
	if id !=1 {
		return nil, []error{errors.New("Not Found")}
	}
	return &sat, nil
}

func(mSatRepo *MockScheduleRepo) StoreSchedule(schedule *entity.Schedule) (*entity.Schedule, []error){
	sat :=schedule
	return sat, nil
}

func (mSatRepo *MockScheduleRepo) InfosInSchedule(schedule *entity.Schedule) ([]entity.Info, []error){
	infos := []entity.Info{entity.InfoMock}
	return infos, nil
}
































