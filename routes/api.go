package api

import (
	userController "laundryApp/app/controllers"

	database "laundryApp/config"

	"github.com/gin-gonic/gin"
)

func CreateMapApiRouters(router *gin.Engine) {

	api := router.Group("/api")
	database.DBConnection()

	api.GET("/user", userController.Index)
	api.POST("/user", userController.Create)
	api.PUT("/user")
	api.DELETE("/user")
}
