package server

import (
	"gin-starter-template/controllers"
	"gin-starter-template/db"
	"gin-starter-template/middlewares"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"

	"github.com/gin-contrib/logger"
)

func Router() *gin.Engine {
	router := gin.New()

	// use middleware
	router.Use(logger.SetLogger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db := db.GetDb()

	store := gormsessions.NewStore(db, true, []byte("secret"))

	router.Use(sessions.Sessions("mysession", store))

	router.Use(middlewares.AuthMiddleware())

	// controllers
	healthController := controllers.HealthController{}
	userController := &controllers.UserController{}
	router.GET("/health", healthController.Status)
	v1 := router.Group("v1")
	{

		user := v1.Group("user")
		user.GET(":id", userController.Retrieve)
		user.POST("save", userController.Create)
	}
	return router
}
