package repositories

import (
	"fmt"
	"goxhactive-final-project/app/models"
	"goxhactive-final-project/config"
)

type SocialMediaRepository interface {
	CreateSosmed(payload models.SocialMedias) (models.SocialMedias, error)
	GetSosmed(userId uint) ([]models.SocialMedias, error)
	UpdateSosmed(payload models.SocialMedias, userId uint) (models.SocialMedias, error)
	DeleteSosmed(payload models.SocialMedias, photoId uint64) (models.SocialMedias, error)
}

func (db *dbConnection) CreateSosmed(payload models.SocialMedias) (models.SocialMedias, error) {

	err := db.connection.Save(&payload).Error
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func NewSosmedRepository() SocialMediaRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *dbConnection) GetSosmed(userId uint) ([]models.SocialMedias, error) {
	var sosmed []models.SocialMedias

	err := db.connection.Where("user_id = ?", userId).Find(&sosmed).Error

	if err != nil {
		return sosmed, err
	}

	return sosmed, nil
}

func (db *dbConnection) UpdateSosmed(payload models.SocialMedias, userId uint) (models.SocialMedias, error) {

	err := db.connection.Model(&payload).Where("id = ?", payload.Id).Where("user_id = ?", userId).Updates(&payload).Error

	if err != nil {
		return payload, err
	}

	return payload, nil
}

func (db *dbConnection) DeleteSosmed(payload models.SocialMedias, sosmedIdInt uint64) (models.SocialMedias, error) {
	var sosmed models.SocialMedias

	err := db.connection.Raw("DELETE FROM social_medias WHERE id = ?", sosmedIdInt).Scan(&sosmed).Error

	fmt.Println(err)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
