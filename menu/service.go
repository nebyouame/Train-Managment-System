package menu

import "TrainSystem/entity"

type ScheduleService interface {
	Schedule()([]entity.Schedule, error)
	StoreSchedule(schedule []entity.Schedule) error
}
