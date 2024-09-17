package router

import (
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")

	userRouter.POST("/register", controllers.UserRegister)
	userRouter.POST("/login", controllers.UserLogin)

	productRouter := r.Group("/products")

	productRouter.Use(middlewares.Authentication())
	productRouter.POST("/", controllers.CreateProduct)

	return r
}
