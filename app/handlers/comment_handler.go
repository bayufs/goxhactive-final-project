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

type CommentHandler struct {
	repository repositories.CommentRepository
}

func CmntHandler() *CommentHandler {
	return &CommentHandler{repositories.NewCommentRepository()}
}

func (commentInstance CommentHandler) CreateComment(ginInstance *gin.Context) {

	var req resources.InputComments

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

	var payload = models.Comments{}

	if err != nil {

		fmt.Println(err.Error())
		ginInstance.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create a comment",
			"status":  http.StatusUnauthorized,
		})

		return
	}

	payload = models.Comments{
		UserID:  userId,
		PhotoId: req.PhotoId,
		Message: req.Message,
	}

	res, err := repository.CreateComment(payload)

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

	response := helpers.JWTResponse("Comment successfully created", http.StatusCreated, payload)
	ginInstance.JSON(http.StatusOK, response)

}

func (commentInstance CommentHandler) GetComment(ginInstance *gin.Context) {

	repository := commentInstance.repository

	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	fmt.Println(userId)

	res, err := repository.GetComment(userId)

	fmt.Println(res, "Results")

	if err != nil {

		ginInstance.JSON(http.StatusNotFound, gin.H{
			"message": "Cant get rows",
			"status":  http.StatusNotFound,
		})

		return
	}

	response := helpers.JWTResponse("All comments ", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)

}

func (commentInstance CommentHandler) UpdateComment(ginInstance *gin.Context) {

	commentId := ginInstance.Param("comment_id")

	commentIdInt, err_strconv := strconv.ParseUint(commentId, 10, 64)
	if err_strconv != nil {
		fmt.Println(err_strconv)
	}
	var req resources.InputCommentsUpdate

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

	var payload = models.Comments{}

	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	payload = models.Comments{
		Id:      uint(commentIdInt),
		Message: req.Message,
	}

	res, err := repository.UpdateComment(payload, userId)

	fmt.Println(res, "Results")

	if err != nil {
		fmt.Println(err)
		ginInstance.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update comment",
			"status":  http.StatusBadRequest,
		})

		return
	}

	response := helpers.JWTResponse("A comment successfully updated ", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)
}

func (commentInstance CommentHandler) DeleteComment(ginInstance *gin.Context) {

	commentId := ginInstance.Param("comment_id")

	commentIdInt, err_strconv := strconv.ParseUint(commentId, 10, 64)
	if err_strconv != nil {
		fmt.Println(err_strconv)
	}

	userData := ginInstance.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["user_id"].(float64))

	repository := commentInstance.repository

	var payload = models.Comments{}

	fmt.Println(userId)

	res, err := repository.DeleteComment(payload, commentIdInt)

	fmt.Println(res, "Results")

	if err != nil {
		fmt.Println(err)
		ginInstance.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete comment",
			"status":  http.StatusBadRequest,
		})

		return
	}

	response := helpers.JWTResponse("Your Comment has been successfully deleted", http.StatusOK, res)
	ginInstance.JSON(http.StatusOK, response)

}
