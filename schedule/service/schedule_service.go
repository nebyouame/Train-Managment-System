package service

import (
	"TrainProject/entity"
	"TrainProject/schedule"
)

type ScheduleService struct {
	scheduleRepo schedule.ScheduleRepository
}

func NewScheduleService(SatRepo schedule.ScheduleRepository) schedule.ScheduleService {
	return &ScheduleService{scheduleRepo:SatRepo}
}

func (ss *ScheduleService) Schedules() ([]entity.Schedule, []error) {
	schedules, errs := ss.scheduleRepo.Schedules()

	if len(errs) > 0 {
		return nil, errs
	}

	return schedules, nil
}

func (ss *ScheduleService) StoreSchedule(schedule *entity.Schedule) (*entity.Schedule, []error){
	sat, errs := ss.scheduleRepo.StoreSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (ss *ScheduleService) Schedule(id uint) (*entity.Schedule, []error) {
	s, errs := ss.scheduleRepo.Schedule(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return s, nil
}

func (ss *ScheduleService) UpdateSchedule(schedule *entity.Schedule) (*entity.Schedule, []error){
	sat, errs := ss.scheduleRepo.UpdateSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}


func (ss *ScheduleService) DeleteSchedule(id uint) (*entity.Schedule, []error) {
	sat, errs := ss.scheduleRepo.DeleteSchedule(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (ss *ScheduleService) InfosInSchedule(schedule *entity.Schedule) ([]entity.Info, []error){
	sts, errs := ss.scheduleRepo.InfosInSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}

	return sts, nil
}






















