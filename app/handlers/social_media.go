package handlers

import (
	"fmt"
	"goxhactive-final-project/app/helpers"
	"goxhactive-final-project/app/models"
	"goxhactive-final-project/app/repositories"
	"goxhactive-final-project/app/resources"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type SocialMediaHandler struct {
	repository repositories.SocialMediaRepository
}

func SosmedHandler() *SocialMediaHandler {
	return &SocialMediaHandler{repositories.NewSosmedRepository()}
}

func (commentInstance SocialMediaHandler) CreateSosmed(ginInstance *gin.Context) {

	var req resources.InputSocialMedias

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
			}

			ginInstance.JSON(http.StatusUnauthorized, gin.H{
				"message": "please check your data request",
				"error":   fmt.Sprint(err_msg),
				"status":  http.StatusBadRequest,
			})
			return
		}
	}

	repository := commentInstance.repository

	var payload = models.SocialMedias{}

	if err != nil {

		fmt.Println(err.Error())
		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a comment",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	payload = models.SocialMedias{
		UserID:         userId,
		SocialMediaUrl: req.SocialMediaUrl,
		Name:           req.Name,
	}

	res, err := repository.CreateSosmed(payload)

	fmt.Println(res, "Results")

	if err != nil {

		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a comment",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	if err != nil {

		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a comment",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	response := helpers.JWTResponse("Social media successfully created", http.StatusCreated, payload)
	ginInstance.JSON(http.StatusOK, response)

}

func (sosmedInstance SocialMediaHandler) GetSosmed(ginInstance *gin.Context) {

	repository := sosmedInstance.repository

	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	fmt.Println(userId)

	res, err := repository.GetSosmed(userId)

	fmt.Println(res, "Results")

	if err != nil {

		ginInstance.JSON(http.StatusNotFound, gin.H{
			"message": "Cant get rows",
			"status":  http.StatusNotFound,
		})

		return
	}

	response := helpers.JWTResponse("All Sosial medias ", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)

}

func (sosmedInstance SocialMediaHandler) UpdateSosmed(ginInstance *gin.Context) {

	sosmedId := ginInstance.Param("socialMediaId")

	sosmedIdInt, err_strconv := strconv.ParseUint(sosmedId, 10, 64)
	if err_strconv != nil {
		fmt.Println(err_strconv)
	}
	var req resources.InputSocialMedias

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

	repository := sosmedInstance.repository

	var payload = models.SocialMedias{}

	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	payload = models.SocialMedias{
		Id:             uint(sosmedIdInt),
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
	}

	res, err := repository.UpdateSosmed(payload, userId)

	fmt.Println(res, "Results")

	if err != nil {
		fmt.Println(err)
		ginInstance.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update social media",
			"status":  http.StatusBadRequest,
		})

		return
	}

	response := helpers.JWTResponse("A social media successfully updated ", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)
}

func (sosmedInstance SocialMediaHandler) DeleteSosmed(ginInstance *gin.Context) {

	sosmedId := ginInstance.Param("socialMediaId")

	sosmedIdInt, err_strconv := strconv.ParseUint(sosmedId, 10, 64)
	if err_strconv != nil {
		fmt.Println(err_strconv)
	}

	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	repository := sosmedInstance.repository

	var payload = models.SocialMedias{}

	fmt.Println(userId)

	res, err := repository.DeleteSosmed(payload, sosmedIdInt)

	fmt.Println(res, "Results")

	if err != nil {
		fmt.Println(err)
		ginInstance.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete social media",
			"status":  http.StatusBadRequest,
		})

		return
	}

	response := helpers.JWTResponse("Your social media has been successfully deleted", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)

}
