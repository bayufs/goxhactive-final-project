package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"goxhactive-final-project/app/helpers"
	"goxhactive-final-project/app/models"
	"goxhactive-final-project/app/repositories"
	"goxhactive-final-project/app/resources"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type PhotoHandler struct {
	repository repositories.PhotoRepository
}

func PhtHandler() *PhotoHandler {
	return &PhotoHandler{repositories.NewPhotoRepository()}
}

func (photoInstance PhotoHandler) CreatePhoto(ginInstance *gin.Context) {

	var req resources.InputPhotos
	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	fmt.Println()

	err := ginInstance.ShouldBind(&req)
	if err != nil {
		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "please check your data request",
			"error":   err,
			"status":  http.StatusBadRequest,
		})
		return
	}

	if err := validate.Struct(&req); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			var err_msg string
			switch fieldErr.Tag() {
			case "required":
				err_msg = fmt.Sprintf("Field : %s is required",
					fieldErr.Field())
			case "gte":
				err_msg = fmt.Sprintf("Field : %s value must be greater than %s charaters",
					fieldErr.Field(), fieldErr.Param())
			}

			ginInstance.JSON(http.StatusUnauthorized, gin.H{
				"message": "please check your data request",
				"error":   fmt.Sprint(err_msg),
				"status":  http.StatusBadRequest,
			})
			return
		}
	}

	repository := photoInstance.repository

	var payload = models.Photos{}

	if err != nil {

		fmt.Println(err.Error())
		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a new photo",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	payload = models.Photos{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   userId,
	}

	res, err := repository.CreatePhoto(payload)

	fmt.Println(res, "Results")

	if err != nil {

		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a new photo",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	if err != nil {

		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a new customer",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	response := helpers.JWTResponse("A new user successfully created ", http.StatusCreated, res)
	ginInstance.JSON(http.StatusOK, response)

}

func (photoInstance PhotoHandler) GetPhoto(ginInstance *gin.Context) {

	repository := photoInstance.repository

	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	fmt.Println(userId)

	res, err := repository.GetPhoto(userId)

	fmt.Println(res, "Results")

	if err != nil {

		ginInstance.JSON(http.StatusNotFound, gin.H{
			"message": "Cant get rows",
			"status":  http.StatusNotFound,
		})

		return
	}

	response := helpers.JWTResponse("All photos ", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)

}

func (photoInstance PhotoHandler) UpdatePhoto(ginInstance *gin.Context) {

	photoId := ginInstance.Param("photo_id")

	photoIdInt, err_strconv := strconv.ParseUint(photoId, 10, 64)
	if err_strconv != nil {
		fmt.Println(err_strconv)
	}
	var req resources.InputPhotos

	err := ginInstance.ShouldBind(&req)
	if err != nil {
		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "please check your data request",
			"error":   err,
			"status":  http.StatusBadRequest,
		})
		return
	}

	if err := validate.Struct(&req); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			var err_msg string
			switch fieldErr.Tag() {
			case "required":
				err_msg = fmt.Sprintf("Field : %s is required",
					fieldErr.Field())
			}

			ginInstance.JSON(http.StatusUnauthorized, gin.H{
				"message": "please check your data request",
				"error":   fmt.Sprint(err_msg),
				"status":  http.StatusBadRequest,
			})
			return
		}
	}

	repository := photoInstance.repository

	var payload = models.Photos{}

	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	payload = models.Photos{
		Id:       photoIdInt,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}

	res, err := repository.UpdatePhoto(payload, userId)

	fmt.Println(res, "Results")

	if err != nil {
		fmt.Println(err)
		ginInstance.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update user",
			"status":  http.StatusBadRequest,
		})

		return
	}

	response := helpers.JWTResponse("A new user successfully created ", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)
}
