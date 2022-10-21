package models

import (
	"gorm.io/gorm"
)

type SocialMedias struct {
	gorm.Model
	Id             uint   `json:"id" gorm:"primary_key" gorm:"column:id"`
	UserID         uint   `json:"user_id"`
	Users          Users  `gorm:"foreignKey:UserID"`
	Name           string `gorm:"foreignkey:photo_id;association_foreignkey:id"`
	SocialMediaUrl string `json:"message" gorm:"column:message"`
}
