package repositories

import (
	"errors"
	"fmt"
	"goxhactive-final-project/app/models"
	"goxhactive-final-project/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	Registration(payload models.Users) (models.Users, error)
	Authorize(email, pass string) (models.Users, error)
	Update(payload models.Users) (models.Users, error)
	Delete(payload models.Users, userId uint) (models.Users, error)
}

func NewAuthenticationRepository() AuthenticationRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *dbConnection) Registration(payload models.Users) (models.Users, error) {

	err := db.connection.Save(&payload).Error
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func (db *dbConnection) Authorize(email, pass string) (models.Users, error) {

	var user models.Users

	err := db.connection.Where("email = ?", email).First(&user)

	fmt.Println(err, "Check if user is exist")

	if errors.Is(err.Error, gorm.ErrRecordNotFound) {

		fmt.Println(err, "User does not exist")

		return user, errors.New("User does not exist")
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if err2 != nil {
		fmt.Println("Hash password result compare : ", err2)
		return user, errors.New("Password does not match")
	}

	return user, nil

}

func (db *dbConnection) Update(payload models.Users) (models.Users, error) {

	err := db.connection.Model(&payload).Where("id = ?", payload.Id).Updates(&payload).Error

	if err != nil {
		return payload, err
	}

	return payload, nil
}

func (db *dbConnection) Delete(payload models.Users, userId uint) (models.Users, error) {
	var user models.Users

	err := db.connection.Raw("DELETE FROM users WHERE id = ?", userId).Scan(&user).Error

	fmt.Println(err)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
