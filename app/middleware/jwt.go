package middleware

import (
	"fmt"
	"goxhactive-final-project/app/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(ginInstance *gin.Context) {
		accessToken := strings.Fields(ginInstance.GetHeader("Authorization"))

		if accessToken[0] != "Bearer" {
			ginInstance.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization is not in Bearer Token format",
				"status":  http.StatusUnauthorized,
			})
			ginInstance.Abort()
			return
		}

		if accessToken[1] == "" {
			ginInstance.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization token is required",
				"status":  http.StatusUnauthorized,
			})
			ginInstance.Abort()
			return
		}
		user, err := helpers.ValidateToken(accessToken[1])
		if err != nil {
			fmt.Println(err)
			ginInstance.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is invalid",
				"status":  http.StatusUnauthorized,
			})
			ginInstance.Abort()
			return
		}
		ginInstance.Set("userData", user)
		ginInstance.Next()
	}
}
