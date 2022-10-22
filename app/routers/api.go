package routers

import (
	"goxhactive-final-project/app/handlers"
	"goxhactive-final-project/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authHandler := handlers.AuthHandler()
	photoHandler := handlers.PhtHandler()
	commentHandler := handlers.CmntHandler()
	sosmedHandler := handlers.SosmedHandler()

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
		photo.DELETE("/:photo_id", photoHandler.DeletePhoto)

	}

	comment := routes.Group("/comments").Use(middleware.ValidateToken())
	{
		comment.POST("/", commentHandler.CreateComment)
		comment.GET("/", commentHandler.GetComment)
		comment.PUT("/:comment_id", commentHandler.UpdateComment)
		comment.DELETE("/:comment_id", commentHandler.DeleteComment)

	}

	sosmed := routes.Group("/socialmedias").Use(middleware.ValidateToken())
	{
		sosmed.POST("/", sosmedHandler.CreateSosmed)
		sosmed.GET("/", sosmedHandler.GetSosmed)
		sosmed.PUT("/:socialMediaId", sosmedHandler.UpdateSosmed)
		sosmed.DELETE("/:socialMediaId", sosmedHandler.DeleteSosmed)

	}

}
