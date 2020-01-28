package entity

import "time"

type Schedule struct {
	ID uint
	TrainSource string `gorm:"type:varchar(255);not null"`
	TrainDestination string
	Image string `gorm:"type:varchar(255)"`
	Infos []Info `gorm:"many2many:info_categories"`
}

type Info struct {
	ID uint
	TrainSource string `gorm:"type:varchar(255);not null"`
	Price float32
	TrainDestination string
	Schedules []Schedule `gorm:"many2many:info_schedules"`
	Image string `gorm:"type:varchar(255)"`
	Paths []Path `gorm:"many2many:info_paths"`
}

type Path struct {
	ID uint
	TrainSource string `gorm:"type:varchar(255)"`
	TrainDestination string
}












