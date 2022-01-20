package main

import (
	"os"

	"github.com/gin-gonic/gin"

	product_controller "microservice/utils/controllers/product_controller"
)

func main() {

	println("env_mysql_pwd", os.Getenv("env_mysql_pwd"))

	r := gin.Default()

	//set api routes
	product_controller.SetApiRoutes(r)

	//start server
	r.Run("localhost:8080")
}
