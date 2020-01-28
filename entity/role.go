package entity

type Role struct {
	ID uint
	Name string `gorm:"type:varchar(255)"`
}
