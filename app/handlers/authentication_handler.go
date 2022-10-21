package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"goxhactive-final-project/app/helpers"
	"goxhactive-final-project/app/models"
	"goxhactive-final-project/app/repositories"
	"goxhactive-final-project/app/resources"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type MyCustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type AuthenticationHandler struct {
	repository repositories.AuthenticationRepository
}

func AuthHandler() *AuthenticationHandler {
	return &AuthenticationHandler{repositories.NewAuthenticationRepository()}
}

func (authInstance AuthenticationHandler) Registration(ginInstance *gin.Context) {

	var req resources.InputUsers

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
			case "email":
				err_msg = fmt.Sprintf("Field : %s is not valid email",
					fieldErr.Field())
			case "min":
				err_msg = fmt.Sprintf("Field : %s value minimal %s charaters",
					fieldErr.Field(), fieldErr.Param())
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

	repository := authInstance.repository

	var payload = models.Users{}

	password := []byte(req.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {

		fmt.Println(err.Error())
		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a new user",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	payload = models.Users{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Age:      req.Age,
	}

	res, err := repository.Registration(payload)

	fmt.Println(res, "Results")

	if err != nil {

		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a new customer",
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

func (authInstance AuthenticationHandler) Login(ginInstance *gin.Context) {
	var req resources.InputUsers
	err := ginInstance.ShouldBind(&req)
	if err != nil {
		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "please check your data request",
			"error":   err,
			"status":  http.StatusBadRequest,
		})
		return
	}

	email := req.Email
	pass := req.Password

	repository := authInstance.repository

	res, err := repository.Authorize(email, pass)

	fmt.Println(res, "Results")

	if err != nil {

		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error credentials, failed to create authentication token",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	access_token, err := helpers.GenerateToken(res)

	if err != nil {
		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create authentication token",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	response := helpers.JWTResponse("success created authorize token", http.StatusOK, access_token)
	ginInstance.JSON(http.StatusOK, response)

}

func (authInstance AuthenticationHandler) Update(ginInstance *gin.Context) {
	var req resources.InputUserUpdate

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
			case "email":
				err_msg = fmt.Sprintf("Field : %s is not valid email",
					fieldErr.Field())
			case "min":
				err_msg = fmt.Sprintf("Field : %s value minimal %s charaters",
					fieldErr.Field(), fieldErr.Param())
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

	repository := authInstance.repository

	var payload = models.Users{}

	userId, err := strconv.ParseUint(ginInstance.Query("userId"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	payload = models.Users{
		Id:        userId,
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now().UTC(),
	}

	res, err := repository.Registration(payload)

	fmt.Println(res, "Results")

	if err != nil {
		fmt.Println(err)
		ginInstance.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update user",
			"status":  http.StatusBadRequest,
		})

		return
	}

	response := helpers.JWTResponse("A new user successfully created ", http.StatusCreated, res)
	ginInstance.JSON(http.StatusOK, response)
}

func (authInstance AuthenticationHandler) Delete(ginInstance *gin.Context) {
	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	repository := authInstance.repository

	var payload = models.Users{}

	fmt.Println(userId)

	res, err := repository.Delete(payload, userId)

	fmt.Println(res, "Results")

	if err != nil {
		fmt.Println(err)
		ginInstance.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete user",
			"status":  http.StatusBadRequest,
		})

		return
	}

	response := helpers.JWTResponse("User successfully deleted", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)

}
