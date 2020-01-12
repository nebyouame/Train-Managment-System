package repository

import (
	"TrainSystem/entity"
	"TrainSystem/train"
	"github.com/jinzhu/gorm"
)

type ScheduleGormRepo struct {
	conn *gorm.DB
}

func NewScheduleGormRepo(db *gorm.DB) train.ScheduleRepository {
	return &ScheduleGormRepo{conn:db}
}

func (sRepo *ScheduleGormRepo) Schedules() ([]entity.Schedule, []error){
	schs := []entity.Schedule{}
	errs := sRepo.conn.Find(&schs).GetErrors()
	if len(errs) > 0{
		return nil, errs
	}
	return schs, errs
}

func (sRepo *ScheduleGormRepo) Schedule(id uint) (*entity.Schedule, []error){
	sch := entity.Schedule{}
	errs := sRepo.conn.First(&sch, id).GetErrors()
	if len(errs) > 0{
		return nil, errs
	}
	return &sch, errs
}

func (sRepo *ScheduleGormRepo) UpdateSchedule(schedule *entity.Schedule)(*entity.Schedule, []error)  {
	sat := schedule
	errs := sRepo.conn.Save(sat).GetErrors()
	if len(errs) > 0{
		return nil, errs
	}
	return sat, errs
}

func (sRepo *ScheduleGormRepo) DeleteSchedule(id uint) (*entity.Schedule, []error){
	sat, errs := sRepo.Schedule(id)
	if len(errs) > 0{
		return nil, errs
	}
	errs = sRepo.conn.Delete(sat, sat.ID).GetErrors()
	if len(errs) > 0{
		return nil, errs
	}
	return sat, errs
}

func (sRepo *ScheduleGormRepo) StoreSchedule(schedule *entity.Schedule)(*entity.Schedule, []error){
	sat := schedule
	errs := sRepo.conn.Create(sat).GetErrors()
	if len(errs) > 0{
		return nil, errs
	}
	return sat, errs
}
func (sRepo *ScheduleGormRepo) ItemInSchedule(schedule *entity.Schedule) ([]entity.Item, []error){
	items := []entity.Item{}
	sat, errs := sRepo.Schedule(schedule.ID)
	if len(errs) > 0{
		return nil, errs
	}
	errs = sRepo.conn.Model(sat).Related(&items, "Items").GetErrors()
	if len(errs) > 0{
		return nil, errs
	}
	return items, errs
}


















































