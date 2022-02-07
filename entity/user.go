package entity

import "time"

type User struct {
	ID        int
	Name      string `gorm:"size:255"`
	Username  string `gorm:"size:255;index:idx_username,unique"`
	Email     string `gorm:"size:255;index:idx_email,unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
