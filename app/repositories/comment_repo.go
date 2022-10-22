package repositories

import (
	"fmt"
	"goxhactive-final-project/app/models"
	"goxhactive-final-project/config"
)

type CommentRepository interface {
	CreateComment(payload models.Comments) (models.Comments, error)
	GetComment(userId uint) ([]models.Comments, error)
	UpdateComment(payload models.Comments, userId uint) (models.Comments, error)
	DeleteComment(payload models.Comments, photoId uint64) (models.Comments, error)
}

func NewCommentRepository() CommentRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *dbConnection) CreateComment(payload models.Comments) (models.Comments, error) {

	err := db.connection.Save(&payload).Error
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func (db *dbConnection) GetComment(userId uint) ([]models.Comments, error) {
	var comment []models.Comments

	err := db.connection.Where("user_id = ?", userId).Find(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (db *dbConnection) UpdateComment(payload models.Comments, userId uint) (models.Comments, error) {

	err := db.connection.Model(&payload).Where("id = ?", payload.Id).Where("user_id = ?", userId).Updates(&payload).Error

	if err != nil {
		return payload, err
	}

	return payload, nil
}

func (db *dbConnection) DeleteComment(payload models.Comments, commentIdInt uint64) (models.Comments, error) {
	var comment models.Comments

	err := db.connection.Raw("DELETE FROM comments WHERE id = ?", commentIdInt).Scan(&comment).Error

	fmt.Println(err)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
