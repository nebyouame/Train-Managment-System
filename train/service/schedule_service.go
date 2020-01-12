package service

import (
	"TrainSystem/entity"
	"TrainSystem/train"
)

type ScheduleService struct {
	scheduleRepo train.ScheduleRepository
}

func NewScheduleService(SatRepo train.ScheduleRepository) train.ScheduleService {
	return &ScheduleService{scheduleRepo: SatRepo}
}

func (ss *ScheduleService) Schedules() ([]entity.Schedule, []error){
	schedules, errs := ss.scheduleRepo.Schedules()

	if len(errs) > 0 {
		return nil, errs
	}
	return schedules, nil
}

func(ss *ScheduleService) StoreSchedule(schedule *entity.Schedule) (*entity.Schedule, []error){
	sat, errs := ss.scheduleRepo.StoreSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}
	return sat,nil
}

func (ss *ScheduleService) Schedule(id uint) (*entity.Schedule, []error){
	c, errs := ss.scheduleRepo.Schedule(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return c, nil
}

func(ss *ScheduleService) UpdateSchedule(schedule *entity.Schedule) (*entity.Schedule, []error){
	sat, errs := ss.scheduleRepo.UpdateSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}
	return sat, nil
}

func (ss *ScheduleService) DeleteSchedule(id uint) (*entity.Schedule, []error){
	sat, errs := ss.scheduleRepo.DeleteSchedule(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return sat, nil
}
func (ss *ScheduleService) ItemsInSchedule(schedule *entity.Schedule) ([]entity.Item, []error){
	sts, errs := ss.scheduleRepo.ItemInSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}
	return sts, nil
}
































