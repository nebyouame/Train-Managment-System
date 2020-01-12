package train

import "TrainSystem/entity"

type ScheduleService interface {
	Schedules()([]entity.Schedule, []error)
	Schedule(id uint) (*entity.Schedule, []error)
	UpdateSchedule(schedule *entity.Schedule) (*entity.Schedule, []error)
	DeleteSchedule(id uint) (*entity.Schedule, []error)
	StoreSchedule(catagory *entity.Schedule) (*entity.Schedule, []error)
	ItemsInSchedule(schedule *entity.Schedule) ([]entity.Item, []error)
}
type ItemService interface {
	Items() ([]entity.Item, []error)
	Item(id uint) (*entity.Item, []error)
	UpdateItem(menu *entity.Item) (*entity.Item, []error)
	DeleteItem(id uint) (*entity.Item, []error)
	StoreItem(item *entity.Item) (*entity.Item, []error)
}
