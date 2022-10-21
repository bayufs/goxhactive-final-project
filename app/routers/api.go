package routers

import (
	"goxhactive-final-project/app/handlers"
	"goxhactive-final-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authHandler := handlers.AuthHandler()
	photoHandler := handlers.PhtHandler()

	routes := r.Group("/api/v1")
	user := routes.Group("/users")
	{
		user.POST("/register", authHandler.Registration)
		user.POST("/login", authHandler.Login)
		user.PUT("/", authHandler.Update).Use(middleware.ValidateToken())
		user.DELETE("/", authHandler.Delete).Use(middleware.ValidateToken())

	}

	photo := routes.Group("/photos").Use(middleware.ValidateToken())
	{
		photo.POST("/", photoHandler.CreatePhoto)
		photo.GET("/", photoHandler.GetPhoto)
		photo.PUT("/:photo_id", photoHandler.UpdatePhoto)

	}

}
