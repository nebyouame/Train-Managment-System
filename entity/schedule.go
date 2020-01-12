package entity

import "time"

type Schedule struct {
	ID uint
	TrainSource string `gorm:"type:varchar(255);not null"`
	TrainDestination string `gorm:"type:varchar(255);not null"`
	Image string `gorm:"type:varchar(255)"`
}

type Role struct {
	ID uint
	Name string `gorm:"type:varchar(255)"`
}

type train struct {
	ID uint
	Name string
	Description string
	image string
}

type Item struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Price       float32
	Description string
	Categories  []Schedule   `gorm:"many2many:item_categories"`
	Image       string       `gorm:"type:varchar(255)"`

}

type Book struct {
	ID       uint
	PlacedAt time.Time
	UserID   uint
	ItemID   uint
	Quantity uint
}


type User struct {
	ID       uint
	UserName string `gorm:"type:varchar(255);not null"`
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	Password string `gorm:"type:varchar(255)"`
	Roles    []Role `gorm:"many2many:user_roles"`
	trains   []train
}

