package train

import "TrainSystem/entity"

type ScheduleService interface {
	Schedules()([]entity.Schedule, error)
	Schedule(id int) (entity.Schedule, error)
	UpdateSchedule(schedule entity.Schedule) error
	DeleteSchedule(id int) error
	StoreSchedule(catagory entity.Schedule) error
}
