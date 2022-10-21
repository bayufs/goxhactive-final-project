package models

import (
	"time"

	"gorm.io/gorm"
)

type Photos struct {
	gorm.Model
	Id        uint64    `json:"id" gorm:"primary_key" gorm:"column:id"`
	Title     string    `json:"title" gorm:"column:title"`
	Caption   string    `json:"caption" gorm:"column:caption"`
	PhotoUrl  string    `json:"photo_url" gorm:"column:photo_url"`
	UserID    uint      `json:"user_id"`
	Users     Users     `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
