package repositories

import (
	"fmt"
	"goxhactive-final-project/app/models"
	"goxhactive-final-project/config"
)

type PhotoRepository interface {
	CreatePhoto(payload models.Photos) (models.Photos, error)
	GetPhoto(userId uint) ([]models.Photos, error)
	UpdatePhoto(payload models.Photos, userId uint) (models.Photos, error)
	DeletePhoto(payload models.Photos, userId uint) (models.Photos, error)
}

func NewPhotoRepository() PhotoRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *dbConnection) CreatePhoto(payload models.Photos) (models.Photos, error) {

	err := db.connection.Save(&payload).Error
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func (db *dbConnection) GetPhoto(userId uint) ([]models.Photos, error) {
	var photo []models.Photos

	err := db.connection.Where("user_id = ?", userId).Find(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (db *dbConnection) UpdatePhoto(payload models.Photos, userId uint) (models.Photos, error) {

	err := db.connection.Model(&payload).Where("id = ?", payload.Id).Where("user_id = ?", userId).Updates(&payload).Error

	if err != nil {
		return payload, err
	}

	return payload, nil
}

func (db *dbConnection) DeletePhoto(payload models.Photos, userId uint) (models.Photos, error) {
	var user models.Users

	err := db.connection.Raw("DELETE FROM users WHERE id = ?", userId).Scan(&user).Error

	fmt.Println(err)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
