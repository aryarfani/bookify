package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `gorm:"type:varchar(191);not null;unique" form:"title" json:"title" binding:"required"`
	Description string    `gorm:"type:varchar(191);not null" form:"description" json:"description" binding:"required"`
	Price       int       `gorm:"type:int(11);not null" form:"price" json:"price" binding:"required"`
	Image       string    `gorm:"type:varchar(191);not null"`
	Status      int       `gorm:"type:int(3);default:0" form:"status" json:"status"`
	StatusText  string    `gorm:"-:all" json:"status_text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      int       `json:"user_id"`
	User        User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
}

type Books []Book

func (book *Book) AfterFind(tx *gorm.DB) (err error) {
	// status_text modifier
	switch book.Status {
	case 0:
		book.StatusText = "Active"
	case 1:
		book.StatusText = "Sale"
	case 44:
		book.StatusText = "Banned"
	}

	// image modifier
	book.Image = "http://localhost:8080/" + book.Image

	return
}
