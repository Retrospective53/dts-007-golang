package router

import (
	"challenge-10/controllers"
	"challenge-10/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", controllers.GetProducts)
		productRouter.GET("/:productId", controllers.GetProductbyId)

		productRouter.Use(middlewares.ProductAuthorization())
		productRouter.Use(middlewares.UserAdminAuthorization())
		productRouter.PUT("/:productId", controllers.UpdateProduct)
		productRouter.DELETE("/:productId", controllers.DeleteProduct)
	}

	return r
}