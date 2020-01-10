package menu

import "TrainSystem/entity"

type ScheduleService interface {
	Schedule(id int)(*entity.Schedule, error)
	StoreSchedule(schedule *entity.Schedule) error
}
