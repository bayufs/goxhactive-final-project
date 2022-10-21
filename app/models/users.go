package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id        uint64    `json:"id" gorm:"primary_key" gorm:"column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	Age       int       `json:"age"  gorm:"age"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
