package entity

type Book struct {
	ID uint
	UserID string
	InfoID string
	Image string `gorm:"type:varchar(255)"`
}

