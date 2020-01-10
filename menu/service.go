package menu

import "github.com/betsegawlemma/web-prog-lab-05-mem/entity"

// CategoryService specifies food menu category related services
type ScheduleService interface {
	Schedule(id int) (*entity.Schedule, error)
	StoreSchedule(category *entity.Schedule) error
}
