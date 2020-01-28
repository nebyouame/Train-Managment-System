package entity

type User struct {
	ID uint
	UserName 	string `gorm:"type:varchar(255);not null"`
	FullName 	string `gorm:"type:varchar(255);not null"`
	Email	string `gorm:"type:varchar(255);not null; unique"`
	Phone 	string `gorm:"type:varchar(255);not null; unique"`
	Password	string `gorm:"type:varchar(255)"`
}

