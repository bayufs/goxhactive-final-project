package models

import (
	"time"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Id        uint      `json:"id" gorm:"primary_key" gorm:"column:id"`
	UserID    uint      `json:"user_id"`
	Users     Users     `gorm:"foreignKey:UserID"`
	PhotoId   uint      `json:"photo_id"`
	Photos    Photos    `gorm:"foreignKey:PhotoId"`
	Message   string    `json:"message" gorm:"column:message"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
