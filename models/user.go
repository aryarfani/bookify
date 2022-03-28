package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `gorm:"type:varchar(191);not null" form:"name" json:"name"`
	Email      string    `gorm:"type:varchar(191);not null:unique" form:"email" json:"email"`
	Password   string    `gorm:"type:varchar(191);not null" form:"password" json:"password"`
	Status     int       `gorm:"type:int(3);default:0" form:"status" json:"status"`
	StatusText string    `gorm:"-:all" json:"status_text"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Users []User

func (user *User) AfterFind(tx *gorm.DB) (err error) {
	// status_text modifier
	switch user.Status {
	case 0:
		user.StatusText = "Active"
	case 1:
		user.StatusText = "Sale"
	case 44:
		user.StatusText = "Banned"
	}

	return
}
