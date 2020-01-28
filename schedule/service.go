package schedule

import "TrainProject/entity"

type ScheduleService interface {
	Schedules() ([]entity.Schedule, []error)
	Schedule(id uint) (*entity.Schedule, []error)
	UpdateSchedule(schedule *entity.Schedule) (*entity.Schedule, []error)
	DeleteSchedule(id uint) (*entity.Schedule, []error)
	StoreSchedule(schedule *entity.Schedule) (*entity.Schedule, []error)
	InfosInSchedule(schedule *entity.Schedule) ([]entity.Info, []error)
}

type RoleService interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

type UserService interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
}

type BookService interface {
	Books() ([]entity.Book, []error)
	Book(id uint) (*entity.Book, []error)
	UpdateBook(book *entity.Book) (*entity.Book, []error)
	DeleteBook(id uint) (*entity.Book, []error)
	StoreBook(book *entity.Book) (*entity.Book, []error)
}










