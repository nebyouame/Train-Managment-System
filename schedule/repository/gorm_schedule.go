package repository

import (
	"TrainProject/entity"
	"TrainProject/schedule"
	"github.com/jinzhu/gorm"
)

type ScheduleGormRepo struct {
	conn *gorm.DB
}

func NewScheduleGormRepo(db *gorm.DB) schedule.ScheduleRepository {
	return &ScheduleGormRepo{conn:db}
}

func (sRepo *ScheduleGormRepo) Schedules() ([]entity.Schedule, []error) {
	schs := []entity.Schedule{}
	errs := sRepo.conn.Find(&schs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return schs, nil
}

func (sRepo *ScheduleGormRepo) Schedule(id uint) (*entity.Schedule, []error) {
	sch := entity.Schedule{}
	errs := sRepo.conn.First(&sch, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return &sch, nil
}

func (sRepo *ScheduleGormRepo) UpdateSchedule(schedule *entity.Schedule) (*entity.Schedule, []error){
	sat := schedule
	errs := sRepo.conn.Save(sat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (sRepo *ScheduleGormRepo) DeleteSchedule(id uint) (*entity.Schedule, []error) {
	sat, errs := sRepo.Schedule(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = sRepo.conn.Delete(sat, sat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (sRepo *ScheduleGormRepo) StoreSchedule(schedule *entity.Schedule) (*entity.Schedule, []error) {
	sat := schedule
	errs := sRepo.conn.Create(sat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (sRepo *ScheduleGormRepo) InfosInSchedule(schedule *entity.Schedule) ([]entity.Info, []error) {
	infos := []entity.Info{}
	sat, errs := sRepo.Schedule(schedule.ID)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = sRepo.conn.Model(sat).Related(&infos, "Infos").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return infos, nil
}

































































