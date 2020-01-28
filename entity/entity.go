package entity

import "time"

type Schedule struct {
	ID uint
	TrainSource string `gorm:"type:varchar(255);not null"`
	TrainDestination string
	Image string `gorm:"type:varchar(255)"`
	Infos []Info `gorm:"many2many:info_categories"`
}












