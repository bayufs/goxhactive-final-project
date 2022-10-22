package main

import (
	"fmt"
	"goxhactive-final-project/app/routers"
	"os"

	"goxhactive-final-project/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	r := gin.Default()

	config.LoadEnv()
	config.MigrateDatabase(db)
	defer config.DisconnectDB(db)
	routers.RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
