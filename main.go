package main

import (
	"github.com/gin-gonic/gin"

	"laundryApp/config"
	RouteApi "laundryApp/routes"
)

func main() {

	route := gin.Default()

	config.DBConnection()

	RouteApi.CreateMapApiRouters(route)

	route.Run()

}
